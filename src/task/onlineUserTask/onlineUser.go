package onlineUserTask

import (
	"rickyWeb/src/constant"
	"rickyWeb/src/utils"
	"rickyWeb/src/utils/zapLog"
	"time"

	"go.uber.org/zap"
)

func StartMemberCheck() {
	go TickForOnlineMember()
}

/* 每三分鐘執行一次 */
func TickForOnlineMember() {
	ticker := time.NewTicker(time.Minute * 3)
	for {
		<-ticker.C
		/* 避免執行操過3分鐘，欲執行項目皆額外呼叫goroutine */
		go OnlineMemberCheckTask()
	}
}

/* 清除已過期用戶 */
func OnlineMemberCheckTask() {
	redisdb := utils.GetRedisDb()
	ssMapCmd := redisdb.HGetAll(constant.LoginOnlineHash)
	if ssMapCmd.Err() != nil {
		zapLog.ErrorW("online memeber task error:", ssMapCmd.Err())
		return
	}
	ssMap := ssMapCmd.Val()
	var delKeyList []string
	for k, v := range ssMap {
		if !utils.CheckExist(k) {
			delKeyList = append(delKeyList, k)
			zapLog.WriteLogInfo("del online member:", zap.String("user", v))
		}
	}
	redisdb.HDel(constant.LoginOnlineHash, delKeyList...)
}
