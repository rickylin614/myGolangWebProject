package onlineUserTask

import (
	"orderbento/src/constant"
	"orderbento/src/utils"
	"orderbento/src/utils/zapLog"
	"time"

	"go.uber.org/zap"
)

func StartMemberCheck() {
	go TickForOnlineMember()
}

func TickForOnlineMember() {
	c := time.Tick(time.Minute * 30)
	for {
		<-c
		go OnlineMemberCheckTask()
	}
}

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
