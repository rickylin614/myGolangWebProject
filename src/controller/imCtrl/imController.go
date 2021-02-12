package imCtrl

import (
	"net/http"
	"rickyWeb/src/models"
	"rickyWeb/src/utils/zapLog"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Manager 所有 websocket 信息
type Manager struct {
	Group                   map[string]map[uint]*Client
	groupCount, clientCount uint
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
}

/* 用戶端資料 */
type Client struct {
	UserID uint
	Name   string
	Group  string
	Conn   *websocket.Conn
	Msg    chan []byte
}

/* 全局 wsManager 管理器 */
var wsManager = Manager{
	Group:       make(map[string]map[uint]*Client),
	Register:    make(chan *Client, 1024),
	UnRegister:  make(chan *Client, 1024),
	groupCount:  0,
	clientCount: 0,
}

func ConnectHandler(ctx *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 驗證
		},
		// 處理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		zapLog.ErrorW("connect error!", err)
		return
	}
	if value, exist := ctx.Get("user"); exist {
		if user, ok := value.(models.User); ok {
			zapLog.WriteLogInfo("connect successs!!")
			cli := &Client{
				UserID: user.ID,
				Name:   user.Name,
				Group:  ctx.DefaultQuery("chat", "chat1"),
				Conn:   conn,
				Msg:    make(chan []byte),
			}
			wsManager.Register <- cli
			go handleMessage(cli)
		}
	}
}

/* 接收每個人發送的訊息 */
func handleMessage(c *Client) {
	defer func() {
		if err := recover(); err != nil {
			wsManager.UnRegister <- c
		}
	}()

	for {
		// err := conn.WriteMessage(websocket.BinaryMessage, []byte("connect success!!"))\
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			zapLog.PanicW("read data error", err)
		}
		go boardCast(message, c)
	}
}

/* 將訊息發送給對應組別所有用戶 */
func boardCast(msg []byte, c *Client) {
	groupMap := wsManager.Group[c.Group]
	jsonStr := `{"m":"` + string(msg) + `","n":"` + c.Name + `"}`
	var byteMsg []byte = []byte(jsonStr)
	for _, cli := range groupMap {
		err := cli.Conn.WriteMessage(websocket.BinaryMessage, byteMsg)
		if err != nil {
			zapLog.PanicW("boardCast error!!", err)
			wsManager.UnRegister <- c
		}
	}
}

/* 初始化 建立統計人數管理goroutine */
func init() {
	go wsManager.ControllRegister()
}

/* 控管連線源資料 */
func (manager *Manager) ControllRegister() {
	for {
		select {
		// 註冊
		case client := <-manager.Register:
			zapLog.WriteLogInfo("client connect! userId:", zap.Uint("userId", client.UserID))
			zapLog.WriteLogInfo("register", zap.Uint("client", client.UserID), zap.String("group", client.Group))

			manager.Lock.Lock()
			if manager.Group[client.Group] == nil {
				manager.Group[client.Group] = make(map[uint]*Client)
				manager.groupCount += 1
			}
			manager.Group[client.Group][client.UserID] = client
			manager.clientCount += 1
			manager.Lock.Unlock()

		// 撤銷
		case client := <-manager.UnRegister:
			zapLog.WriteLogInfo("unregister", zap.Uint("client", client.UserID), zap.String("group", client.Group))
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Group]; ok {
				if _, ok := manager.Group[client.Group][client.UserID]; ok {
					close(client.Msg)
					delete(manager.Group[client.Group], client.UserID)
					manager.clientCount -= 1
					if len(manager.Group[client.Group]) == 0 {
						zapLog.WriteLogInfo("unregister", zap.String("group", client.Group))
						delete(manager.Group, client.Group)
						manager.groupCount -= 1
					}
				}
			}
			manager.Lock.Unlock()
		}
	}
}

// func Register() {

// }
