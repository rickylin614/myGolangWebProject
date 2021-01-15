package contoller

import (
	"fmt"
	"net/http"
	"time"

	"orderbento/src/constant"
	"orderbento/src/dao"
	"orderbento/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userReq struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

func Register(ctx *gin.Context) {
	var data userReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := make(gin.H)
	user := dao.QueryUserByName(data.Name)
	fmt.Println(user)
	if user.ID != 0 {
		resp["msg"] = "已註冊的帳號"
		resp["code"] = "error"
	} else {
		user.Name = data.Name
		user.Pwd = data.Password
		user.Insert()
		resp["msg"] = "註冊成功"
	}
	ctx.JSON(http.StatusOK, resp)
}

func Login(ctx *gin.Context) {
	var data userReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := dao.QueryUserByName(data.Name)
	if user.ID != 0 {
		user.SessionId = uuid.New().String()
		redisdb := utils.GetRedisDb()
		redisdb.Set(constant.LoginKey+user.SessionId, time.Now().Nanosecond(), time.Hour*3)
		ctx.SetCookie("sessionId", user.SessionId, int(time.Hour*3), "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "登入成功",
			"code": http.StatusOK,
		})
		user.UpdateLoginTime()
	}
}

func Logout(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("sessionId")
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetCookie("sessionId", "", -1, "/", "", false, true) //時間設為-1 刪除cookie
	redisdb := utils.GetRedisDb()
	key := constant.LoginKey + sessionId
	cmd := redisdb.Del(key)
	if cmd.Err() != nil {
		fmt.Println(cmd.Err())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登出成功",
	})
}
