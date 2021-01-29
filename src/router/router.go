package router

import (
	"fmt"
	"orderbento/src/contoller"

	"github.com/gin-gonic/gin"
)

func RouterSetting(router *gin.Engine) {
	mainRouter := router.Group("/go/_ajax")

	//main setting
	{
		mainRouter.GET("/linkCheck", contoller.LinkCheck)
	}

	userGroup := mainRouter.Group("/user")
	{ //set user handler here
		userGroup.POST("/register", contoller.Register)
		userGroup.POST("/login", contoller.Login)
		userGroup.GET("/logout", contoller.Logout)
		userGroup.POST("/queryUser", contoller.QueryUser)
		userGroup.POST("/loginRecord", contoller.LoginRecord)
	}

	storeGroup := mainRouter.Group("/store")
	{ //set user handler here
		storeGroup.POST("/insert", contoller.InsertStore)
		storeGroup.POST("/update", contoller.UpdateStore)
		storeGroup.POST("/delete", contoller.DeleteStore)
		storeGroup.POST("/queryStoreById", contoller.QueryStoreById)
		storeGroup.POST("/queryStore", contoller.QueryStore)

	}

	menuGroup := mainRouter.Group("/menu")
	{ //set menu handler here
		fmt.Println(menuGroup)
	}

}
