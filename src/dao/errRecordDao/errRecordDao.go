package errRecordDao

import (
	"rickyWeb/src/dao"
	"rickyWeb/src/models"

	"github.com/jinzhu/gorm"
)

type ErrRecord models.ErrRecord

func (r ErrRecord) TableName() string {
	return "err_record"
}

func db() *gorm.DB {
	return dao.GetDB()
}

func Insert(r *ErrRecord) {
	db().Create(r)
}
