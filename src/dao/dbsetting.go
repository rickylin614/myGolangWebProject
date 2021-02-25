package dao

import (
	"log"
	"os"
	"rickyWeb/src/utils/viperUtils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dbtype := "mysql"
	dsn := viperUtils.GetCommonParams("sqlPath")
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
	// 設定顯示dao log
	db.LogMode(true)
}

func SetLogFile(out *os.File) {
	logger := log.New(out, "\r\n", log.LstdFlags)
	db.SetLogger(logger)
}

// 避免任何程序對初始化後的DB做設定修改，連線後"只能"對db設定做"讀取"
// return a clone db
func GetDB() *gorm.DB {
	db := db
	return db
}
