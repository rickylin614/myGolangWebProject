package contoller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Suc = "success"
	Err = "error"
)

func LinkCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "link success!",
		"code": Suc,
	})
}

func NoSetting(ctx *gin.Context) {
	data := gin.H{
		"msg":  "undefined path!",
		"code": Err,
	}
	ctx.JSON(http.StatusNotFound, data)
}
