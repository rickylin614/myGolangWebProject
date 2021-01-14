package server

import (
	"orderbento/src/contoller"
	"orderbento/src/middleware"

	"github.com/gin-gonic/gin"
)

//default setting

func GinInit() *gin.Engine {
	router := gin.Default()
	router.NoRoute(contoller.NoSetting)
	router.NoMethod(contoller.NoSetting)

	//中間件設定
	{
		router.Use(middleware.LoginCheck) //登入中間件
	}
	return router
}
