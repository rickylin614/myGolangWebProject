package middleware

import (
	"net/http"
	"rickyWeb/src/models"
	"rickyWeb/src/service/errRecordService"
	"rickyWeb/src/utils"
	"rickyWeb/src/utils/zapLog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ErrorData struct {
	request  http.Request
	userName string
	userId   uint
	errmsg   string
}

var errChan chan ErrorData

func init() {
	errChan = make(chan ErrorData, 1024)
	go errChanWorker()
}

func Common(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			userName := "unknown"
			userId := uint(0)
			if value, exist := ctx.Get("user"); exist {
				if user, ok := value.(models.User); ok {
					userName = user.Name
					userId = user.ID
				}
			}
			if e, ok := err.(error); ok {
				zapLog.WriteLogError("error! ", zap.String("user:", userName), zap.Error(e)) //將錯誤訊息記錄到log文件
				errChan <- ErrorData{                                                        //將錯誤訊息透過channel紀錄到DB 方便查詢
					request:  *ctx.Request,
					userName: userName,
					userId:   userId,
					errmsg:   e.Error(),
				}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg":  "請求發生錯誤! 請稍後重新再試",
				"code": "error",
			})
		}
	}()
	start := time.Now()
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	ip := utils.GetRealIp(ctx.Request)

	ctx.Next()

	//計算消耗時間
	t := time.Since(start).Milliseconds()
	strT := strconv.FormatInt(t, 10) + "ms"
	if t == 0 {
		t := time.Since(start).Microseconds()
		strT = strconv.FormatInt(t, 10) + "ns"
	}

	zapLog.WriteLogInfo(
		"[gin]",
		zap.String("time", strT),
		zap.Int("status code", ctx.Writer.Status()),
		zap.String("ip", ip),
		zap.String("method", method),
		zap.String("path", path),
	)
}

func errChanWorker() {
	for {
		data := <-errChan
		errRecordService.Insert(&data.request, data.userId, data.userName, data.errmsg)
	}
}
