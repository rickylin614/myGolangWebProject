package loginRecordService

import (
	"fmt"
	"net/http"
	"orderbento/src/dao/loginRecordDao"
	"orderbento/src/dao/userDao"
	"orderbento/src/utils"
	"time"
)

var recordChan chan loginRecordDao.LoginRecord

/* 初始化channel */
func init() {
	recordChan = make(chan loginRecordDao.LoginRecord, 1024)
	go catchRecord()
}

/* 查詢登入紀錄 */
func Index(params map[string]interface{}) (records []loginRecordDao.LoginRecord, count int) {
	var newParams map[string]interface{}
	pageNo, pageSize := utils.GetPage(params)
	utils.CopyParams(params, newParams, "name", "userId", "ip", "loginState")

	count = loginRecordDao.Count(newParams)
	if count != 0 {
		records = loginRecordDao.Query(pageNo, pageSize, newParams)
	}
	return
}

/* 登入紀錄送入chan並回傳 */
func Insert(req *http.Request, user userDao.User, loginState int) {
	headerStr := fmt.Sprint(req.Header)
	record := loginRecordDao.LoginRecord{
		Name:       user.Name,
		UserId:     user.ID,
		LoginTime:  time.Now(),
		UserAgent:  req.UserAgent(),
		Ip:         utils.GetRealIp(req),
		Header:     headerStr,
		LoginState: loginState,
	}
	recordChan <- record
}

/* 接收登入紀錄並寫入DB */
func catchRecord() {
	for {
		record := <-recordChan
		loginRecordDao.Insert(&record)
	}
}
