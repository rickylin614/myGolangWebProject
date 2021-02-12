package main

import (
	_ "github.com/go-sql-driver/mysql"

	"orderbento/src/router"
	"orderbento/src/server"
	"orderbento/src/task"
)

func main() {
	task.Start() //啟動排程

	engine := server.GinInit()   //初始化gin設定
	router.RouterSetting(engine) //路由設定

	err := engine.Run(":8081")
	if err != nil {
		panic(err)
	}
}
