package router

import (
	"fmt"
	"orderbento/src/controller"
	"orderbento/src/controller/imCtrl"

	"github.com/gin-gonic/gin"
)

func RouterSetting(router *gin.Engine) {
	mainRouter := router.Group("/go/_ajax")

	//main setting
	{
		mainRouter.GET("/linkCheck", controller.LinkCheck)
		mainRouter.GET("/chat", imCtrl.ConnectHandler)
	}

	userGroup := mainRouter.Group("/user")
	{ //set user handler here
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/login", controller.Login)
		userGroup.GET("/logout", controller.Logout)
		userGroup.POST("/queryUser", controller.QueryUser)
		userGroup.POST("/loginRecord", controller.LoginRecord)
		userGroup.POST("/onlineMemberList", controller.OnlineMemberList)
		userGroup.POST("/onlineMemberKick", controller.OnlineMemberKick)
	}

	storeGroup := mainRouter.Group("/store")
	{ //set user handler here
		storeGroup.POST("/insert", controller.InsertStore)
		storeGroup.POST("/update", controller.UpdateStore)
		storeGroup.POST("/delete", controller.DeleteStore)
		storeGroup.POST("/queryStoreById", controller.QueryStoreById)
		storeGroup.POST("/queryStore", controller.QueryStore)

	}

	menuGroup := mainRouter.Group("/menu")
	{ //set menu handler here
		fmt.Println(menuGroup)
	}

}
