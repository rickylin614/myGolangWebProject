package errRecordService

import (
	"net/http"
	"rickyWeb/src/dao/errRecordDao"
	"rickyWeb/src/utils"
)

/* 登入紀錄送入chan並回傳 */
func Insert(req *http.Request, userId uint, userName, err string) {
	record := errRecordDao.ErrRecord{
		Name:      userName,
		UserId:    userId,
		ErrMsg:    req.URL.Path + ":" + err,
		UserAgent: req.UserAgent(),
		Ip:        utils.GetRealIp(req),
	}
	errRecordDao.Insert(&record)
}
