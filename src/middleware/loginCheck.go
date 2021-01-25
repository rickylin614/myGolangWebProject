package middleware

import (
	"fmt"
	"net/http"
	"orderbento/src/constant"
	"orderbento/src/utils"
	"orderbento/src/utils/zapLog"

	"github.com/gin-gonic/gin"
)

var out gin.H = gin.H{
	"code": "notLogin",
	"msg":  "尚未登入",
}

func LoginCheck(ctx *gin.Context) {
	//before
	if ctx.Request.URL.Path == "/go/_ajax/user/login" ||
		ctx.Request.URL.Path == "/go/_ajax/user/register" {
		return
	}
	data, err := ctx.Cookie("sessionId")

	if err != nil {
		zapLog.ErrorW("login check error!:", err)
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()
		return
	}
	redisdb := utils.GetRedisDb()
	cmd := redisdb.Get(constant.LoginKey + data)
	if cmd.Err() != nil || cmd.Val() == "" {
		fmt.Printf("err: %v , value %v\n", cmd.Err(), cmd.Val())
		ctx.JSON(http.StatusOK, out)
		ctx.Abort()
		return
	} /* else {
		fmt.Printf("check login success! %v\n", cmd.Val())
	} */
	ctx.Next()
}
