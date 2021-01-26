package contoller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"orderbento/src/constant"
	"orderbento/src/dao/loginRecordDao"
	"orderbento/src/dao/userDao"
	"orderbento/src/models"
	"orderbento/src/service/loginRecordService"
	"orderbento/src/service/userService"
	"orderbento/src/utils"
	"orderbento/src/utils/zapLog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userReq struct {
	Name     string `binding:"required"`
	Password string `binding:"required"`
}

/* 會員註冊 */
func Register(ctx *gin.Context) {
	var data userReq
	err := ctx.Bind(&data)
	if err != nil {
		zapLog.ErrorW("register error!:", err)
		return
	}
	resp := make(gin.H)
	user := userService.QueryUserByName(data.Name)
	zapLog.WriteLogInfo("user register", zap.String("name", user.Name))
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

/* 登入 */
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
		now := time.Now()
		user.LoginTime = &now
		userJson, err := json.Marshal(user)
		if err != nil {
			zapLog.ErrorW("json error!", err)
			return
		}
		redisdb.Set(constant.LoginKey+user.SessionId, userJson, time.Hour*3)
		ctx.SetCookie("sessionId", user.SessionId, int(time.Hour*3), "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "登入成功",
			"code": http.StatusOK,
		})
		userService.UpdateLoginTime(user)
		loginRecordService.Insert(ctx.Request, user, constant.Login)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "查無使用者帳號",
			"code": "error",
		})
		return
	}
}

/* 登出 */
func Logout(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("sessionId")
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetCookie("sessionId", "", -1, "/", "", false, true) //時間設為-1 刪除cookie
	key := constant.LoginKey + sessionId
	jsonStr := utils.GetDel(key)
	var user userDao.User
	err = json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		zapLog.ErrorW("logout Unmarshal error!:", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登出成功",
	})
	loginRecordService.Insert(ctx.Request, user, constant.LogOut)
}

/* 查詢用戶頁面 */
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

/* 組合查詢用戶回傳資料 */
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

/* 登入紀錄頁面 */
func LoginRecord(ctx *gin.Context) {
	var params gin.H
	err := ctx.Bind(&params)
	if err != nil {
		zapLog.ErrorW("LoginRecord error!", err)
		return
	}
	records, count := loginRecordService.Index(params)
	data := composeLoginRecordResp(records)
	ctx.JSON(http.StatusOK, gin.H{
		"data":      data,
		"dataCount": count,
		"msg":       Suc,
	})
}

/* 組合登入紀錄回傳資料 */
func composeLoginRecordResp(records []loginRecordDao.LoginRecord) []models.LoginRecordResponse {
	resps := make([]models.LoginRecordResponse, 0, len(records))
	for _, record := range records {
		var state string
		if record.LoginState == 0 {
			state = "登入"
		} else {
			state = "登出"
		}
		resp := models.LoginRecordResponse{
			Id:         record.Id,
			Name:       record.Name,
			UserId:     record.Id,
			LoginTime:  utils.TimeToString(&record.LoginTime),
			Ip:         record.Ip,
			LoginState: state,
		}
		resps = append(resps, resp)
	}
	return resps
}
