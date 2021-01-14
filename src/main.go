package main

import (
	_ "github.com/go-sql-driver/mysql"

	"orderbento/src/router"
	"orderbento/src/server"
)

func main() {
	engine := server.GinInit()
	router.RouterSetting(engine) //路由設定

	err := engine.Run(":8081")
	if err != nil {
		panic(err)
	}
}
