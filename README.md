

# Golang Web專案

- 採前後分離
- 對應前端Vue : <https://github.com/rickylin614/myGolangWebProject>

# 前端:
	使用前端框架: vue.js(v2.5.2)
-	套件版本:
	>axios: v0.21.1,<br>
	element-ui: v2.14.1,<br>
    	vue-router: v3.0.1,<br>
    	vuex: v3.6.0,<br>
	net: v1.0.2,
	
# 後端:
	使用語言:golang
-	依賴:
	>gin-gonic/gin v1.6.3  // 網頁框架<br>
	go-sql-driver/mysql v1.5.0 // SQL DRIVER<br>
	jinzhu/gorm v1.9.16 // ORM套件<br>
	google/uuid v1.1.4 // uuid套件<br>
	uber-go/zap v1.16.0 // 日誌記錄套件<br>
	github.com/spf13/viper v1.7.1 // 設定檔存讀套件<br>
	
-   簡易流程
![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/00.簡易架構流程.png) 

## 其他技術應用:
-
	>ngnix : http跳轉使用/IP限制使用<br>
	redis server : 登入緩存使用<br>
	mysql server : 系統資料保存<br>

## 依賴包說明
### 1. gin-gonic

- 一種go的http web framework
- 使用gin來快速配置路由以及中間件(達成類似JAVA AOP)
- 在benchmark中，效能高於其他frame

### 2. uber-go/zap

- uber團隊的zap包用於日誌紀錄
- 高性能的日誌紀錄工具，效能高於其餘日誌紀錄
- 具有報錯顯示完整路徑並不會強制中斷程序的error錯誤日誌類型

### 3. viper

- viper為設定檔管理工具
- 支持JSON，TOML，YAML，HCL，envfile，Properties等等檔案
- 支持從環境變數中存取
- 支持熱讀取

### 4. gorm 

- ORM工具，方便從資料庫中存寫數據。
- 可快速的使用預設CRUD，使用方法輕鬆易懂
- 可自定義Logger，且默認Logger已有足夠的資料
- 支持事務回滾

### 5. go-redis

- 緩存存儲工具
- 使用於登入緩存、在線會員緩存、資料鎖

### 6. google/uuid

- 快速產生不會重複的唯一值
- 目前作用存取用戶識別(產生session id給予用戶)

### 7. gorilla/websocket

- 用於快速部屬長連線通訊協定
- 用於即時回傳訊息給用戶
- 用於即時給前端新增好友通知、新群組通知等等提示訊息


## 功能介紹

### 1. 註冊

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/02.註冊頁面.jpg) 

- 用戶註冊 簡易的帳號密碼即可註冊 未來考慮增加驗證碼增加複雜度以及防機器人註冊
> /src/controller/userController.go Register

```
// 取得前端請求
var data userReq
err := ctx.Bind(&data)	
if err != nil {
	zapLog.ErrorW("register error!:", err)
	return
}
resp := make(gin.H)
user := userService.QueryUserByName(data.Name)
zapLog.WriteLogInfo("user register", zap.String("name", user.Name))

// 判斷帳號是否可註冊，並回傳json訊息給予前端
if user.ID != 0 {
	resp["msg"] = "已註冊的帳號"
	resp["code"] = "error"
} else {
	user.Name = data.Name
	user.Pwd = data.Password
	userService.Insert(user)
	resp["msg"] = "註冊成功"
}
ctx.JSON(http.StatusOK, resp)
```

### 2. 登入

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/01.登入頁面.jpg) 

- 登入主要為驗證使用者輸入資料正確，以及存入REDIS和用戶端cookie
> /src/controller/userController.go Login
```
...
user := userService.QueryUserByName(data.Name)  // 由user.userService查詢DB
if user.ID != 0 { 								// 若不存在則登入失敗
	if user.Pwd != data.Password { /* 密碼檢查 start */
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "密碼錯誤",
			"code": "error",
		})
		return
	}

	user.SessionId = uuid.New().String()		//產生新的UUID
	... /* 解析部分資料 */
	redisdb.Set(constant.LoginKey+user.SessionId, userJson, time.Hour*3)				//資料放入redis緩存 存活時間三小時(配合3.驗證登入使用)
	redisdb.HSet(constant.LoginOnlineHash, constant.LoginKey+user.SessionId, userJson)	//資料放入在線會員清單(配合6.在線會員使用)
	ctx.SetCookie("sessionId", user.SessionId, int(time.Hour*3), "/", "", false, true)	//資料放入瀏覽器Cookie
	...
	userService.UpdateLoginTime(user)	//將登入資訊寫回DB
	loginRecordService.Insert(ctx.Request, user, constant.Login) //將登入紀錄存入DB(內部由goroutine達成寫入DB減少耗時)
...
```

### 3. 中間件驗證登入狀態

- 設定於所有登入後API，於gin中設定好的中間件，驗證是否登入狀態
> /src/middleware/loginCheck.go LoginCheck

```
// src/server/server.go
...
//中間件設定
{
	router.Use(middleware.Common)     //登入中間件
	router.Use(middleware.LoginCheck) //登入中間件 <--由此設定router中間件
}
...
```

```
//非登入時的輸出狀態
var out gin.H = gin.H{
	"code": "notLogin",
	"msg":  "尚未登入",
}
...
data, err := ctx.Cookie("sessionId")	//讀取用戶cookie

if err != nil {
	zapLog.ErrorW("login check error!:", err)
	ctx.JSON(http.StatusOK, out)
	ctx.Abort()	//Abort: 不繼續執行其餘handler
	return
}

// 存取redsi 若已經有資料且可轉models.User則Pass
redisdb := utils.GetRedisDb()
cmd := redisdb.Get(constant.LoginKey + data)
if cmd.Err() != nil || cmd.Val() == "" {	//無登入資訊判定為非登入
	fmt.Printf("err: %v , value %v\n", cmd.Err(), cmd.Val())
	ctx.JSON(http.StatusOK, out)
	ctx.Abort()
	return
} else {
	redisdb.Expire(constant.LoginKey, time.Hour*3) 	//若驗證通過則延長登入時效三小時
	var user models.User
	err := json.Unmarshal([]byte(cmd.Val()), &user)	//從redis中取得的user資料解析為物件
	if err != nil {
		zapLog.ErrorW("login check err!", err)
		ctx.Abort()
		return
	}
	ctx.Set("user", user)  // 存入gin.context資料 此流程後續任何handler皆可取用使用者登入資訊
}
ctx.Next() //繼續執行其餘handler
```

- 若驗證不通過 前端接收回傳json判斷code為notLogin時 則出現彈窗並跳轉到登入頁面

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/04.未登入提示.jpg)

### 4. 用戶頁面

- 可查看所有用戶帳號以及註冊時間、最後登入時間

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/03.用戶管理頁面.jpg)
> /src/controller/userController.go QueryUser

```
func QueryUser(ctx *gin.Context) {
	var req map[string]interface{}
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	users, count := userService.QueryUser(req)	//從DB查詢資料(含分頁參數)
	userResps := composeUserResp(users)	//組合回傳資料
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "查詢成功",
		"data":      &userResps,
		"dataCount": count,
	})
}

...

/* 組合查詢用戶回傳資料 (過濾掉不必要的參數 如密碼，並且自定義轉換時間參數)*/
func composeUserResp(us []userDao.User) []models.UserResponse {
	urs := make([]models.UserResponse, 0, len(us))
	var ursp models.UserResponse
	for _, user := range us {
		ursp = models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			LoginTime: utils.TimeToString(user.LoginTime),
			CreatedAt: utils.TimeToString(&user.CreatedAt),
		}
		urs = append(urs, ursp)
	}
	return urs
}
```

### 5. 登入記錄

- 查詢使用者登入紀錄

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/05.登入紀錄.jpg)
> /src/controller/userController.go LoginRecord

```
var params gin.H
err := ctx.Bind(&params)
if err != nil {
	zapLog.ErrorW("LoginRecord error!", err)
	return
}
records, count := loginRecordService.Index(params) // 條件帶給service層查詢
data := composeLoginRecordResp(records)	//組回傳資料
ctx.JSON(http.StatusOK, gin.H{
	"data":      data,
	"dataCount": count,
	"msg":       Suc,
})

```

### 6. 在線會員

- 每當用戶登入時，存入redis hash，登出時踢出，並有排程每三分鐘執行踢出過期的登入。
- 該頁面設有踢出功能，踢出後該會員必須重新登入才可進行操作。
> /src/controller/userController.go OnlineMemberList
> /src/controller/userController.go OnlineMemberKick
![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/06.在線會員.jpg)

- 查詢程式碼

```
redisdb := utils.GetRedisDb()
ssMapCmd := redisdb.HGetAll(constant.LoginOnlineHash) //從redis中取得map
...
var userList []userDao.User
for _, v := range ssMapCmd.Val() { //遍歷map並解析數據
	var user userDao.User
	err := json.Unmarshal([]byte(v), &user)
	if err == nil {
		userList = append(userList, user)
	}
}
resData := composeUserResp(userList) //組返回資料結構
ctx.JSON(http.StatusOK, gin.H{
	"data": resData,
	"msg":  Suc,
})
```

- 排程(三分鐘執行過期踢出)
```
/* 每三分鐘執行一次 */
func TickForOnlineMember() {
	ticker := time.NewTicker(time.Minute * 3)
	for {
		<-ticker.C
		/* 避免執行超過3分鐘，欲執行項目皆額外呼叫goroutine */
		go OnlineMemberCheckTask()
	}
}

/* 清除已過期用戶 */
func OnlineMemberCheckTask() {
	redisdb := utils.GetRedisDb()
	ssMapCmd := redisdb.HGetAll(constant.LoginOnlineHash)	// 讀取在線列表
	if ssMapCmd.Err() != nil {
		zapLog.ErrorW("online memeber task error:", ssMapCmd.Err())
		return
	}
	ssMap := ssMapCmd.Val()
	var delKeyList []string
	for k, v := range ssMap {	// 遍歷列表
		if !utils.CheckExist(k) {	// 檢查是否在線
			delKeyList = append(delKeyList, k)
			zapLog.WriteLogInfo("del online member:", zap.String("user", v))
		}
	}
	redisdb.HDel(constant.LoginOnlineHash, delKeyList...)	//統一移除不在線，減少連線次數
}	
```

### 7. 即時通訊

- 可進入聊天室房間，進行不同房間之間的溝通對話，採用websocket長連線，前端可即時收到新訊息
> /src/controller/imCtrl/imController.go

![](https://github.com/rickylin614/myGolangWebProject/raw/master/resource/image/07.聊天室演示.jpg)

利用goroutine監聽每個成功連接的用戶，會在ControllRegister裡接收註冊/註銷的用戶，寫到全域管理方法

- 聊天室簡易演示影片 <https://youtu.be/gJDfWJpQvNo>

