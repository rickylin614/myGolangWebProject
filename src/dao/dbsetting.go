package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dbtype := "mysql"
	dsn := "root:qwe123@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true"
	db, err = gorm.Open(dbtype, dsn)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
}

// return a clone db
func GetDB() *gorm.DB {
	db := db
	return db
}
