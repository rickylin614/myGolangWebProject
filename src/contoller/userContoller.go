package contoller

import (
	"fmt"
	"net/http"
	"time"

	"orderbento/src/constant"
	"orderbento/src/dao/userDao"
	"orderbento/src/models"
	"orderbento/src/service/userService"
	"orderbento/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//D:\orderSystem\backend_order_bento\src\service\userService\userService.go

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
	user := userService.QueryUserByName(data.Name)
	fmt.Println(user)
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
}

func Login(ctx *gin.Context) {
	var data userReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := userService.QueryUserByName(data.Name)
	if user.ID != 0 {
		/* 密碼檢查 start */
		if user.Pwd != data.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "密碼錯誤",
				"code": "error",
			})
			return
		}
		/* 密碼檢查 end */
		user.SessionId = uuid.New().String()
		redisdb := utils.GetRedisDb()
		redisdb.Set(constant.LoginKey+user.SessionId, time.Now().Nanosecond(), time.Hour*3)
		ctx.SetCookie("sessionId", user.SessionId, int(time.Hour*3), "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "登入成功",
			"code": http.StatusOK,
		})
		userService.UpdateLoginTime(user)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "查無使用者帳號",
			"code": "error",
		})
		return
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

func QueryUser(ctx *gin.Context) {
	var req map[string]interface{}
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	users, count := userService.QueryUser(req)
	userResps := composeUserResp(users)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "查詢成功",
		"data":      &userResps,
		"dataCount": count,
	})
}

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
