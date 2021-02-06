package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orderbento/src/constant"
	"orderbento/src/models"
	"orderbento/src/utils"
	"orderbento/src/utils/zapLog"
	"time"

	"github.com/gin-gonic/gin"
)

var out gin.H = gin.H{
	"code": "notLogin",
	"msg":  "尚未登入",
}

func LoginCheck(ctx *gin.Context) {
	// 不驗證登入/註冊API
	if ctx.Request.URL.Path == "/go/_ajax/user/login" ||
		ctx.Request.URL.Path == "/go/_ajax/user/register" {
		return
	}
	data, err := ctx.Cookie("sessionId") //讀取用戶cookie

	if err != nil {
		zapLog.ErrorW("login check error!:", err)
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()
		return
	}

	// 存取redsi 若已經有資料且可轉models.User則Pass
	redisdb := utils.GetRedisDb()
	cmd := redisdb.Get(constant.LoginKey + data)
	if cmd.Err() != nil || cmd.Val() == "" {
		fmt.Printf("err: %v , value %v\n", cmd.Err(), cmd.Val())
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()
		return
	} else {
		redisdb.Expire(constant.LoginKey, time.Hour*3)
		var user models.User
		err := json.Unmarshal([]byte(cmd.Val()), &user)
		if err != nil {
			zapLog.ErrorW("login check err!", err)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
	}
	ctx.Next()
}
