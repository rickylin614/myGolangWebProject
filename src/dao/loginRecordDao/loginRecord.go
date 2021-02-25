package loginRecordDao

import (
	"rickyWeb/src/dao"
	"rickyWeb/src/models"

	"github.com/jinzhu/gorm"
)

type LoginRecord models.LoginRecord

func (r LoginRecord) TableName() string {
	return "login_record"
}

func db() *gorm.DB {
	return dao.GetDB()
}

func Insert(r *LoginRecord) {
	db().Create(r)
}

func Query(pageNo, pageSize int, data map[string]interface{}) []LoginRecord {
	records := []LoginRecord{}
	offset := (pageNo - 1) * pageSize
	db().Offset(offset).Limit(pageSize).Find(&records, data)
	return records
}

func Count(data map[string]interface{}) (count int) {
	db().Model(&LoginRecord{}).Where(data).Count(&count)
	return
}
