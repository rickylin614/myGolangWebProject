package server

import (
	"orderbento/src/contoller"
	"orderbento/src/middleware"

	"github.com/gin-gonic/gin"
)

//default setting

func GinInit() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) //設定模式
	// 設定日誌輸出位置
	router := gin.Default()
	router.NoRoute(contoller.NoSetting)
	router.NoMethod(contoller.NoSetting)

	//中間件設定
	{
		router.Use(middleware.Common)     //登入中間件
		router.Use(middleware.LoginCheck) //登入中間件
	}
	return router
}
