package server

import (
	"os"
	"rickyWeb/src/controller"
	"rickyWeb/src/dao"
	"rickyWeb/src/middleware"
	"rickyWeb/src/utils/viperUtils"

	"github.com/gin-gonic/gin"
)

//default setting

func GinInit() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) //設定模式

	// 設定日誌輸出位置
	if gin.Mode() == gin.ReleaseMode {
		file, _ := os.Create(viperUtils.GetLogPath()) //日誌檔案位置
		gin.DefaultWriter = file                      //設定日誌輸出模式
		gin.DefaultErrorWriter = file                 //設定錯誤日誌
		dao.SetLogFile(file)                          //設定gorm日誌輸出
	}

	router := gin.New()
	router.NoRoute(controller.NoSetting)
	router.NoMethod(controller.NoSetting)

	//中間件設定
	{
		router.Use(middleware.Common)     //基本中間件
		router.Use(middleware.LoginCheck) //登入中間件
	}
	return router
}
