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
	data := map[string]string{
		"msg":  "link success!",
		"code": Suc,
	}
	ctx.JSON(http.StatusOK, data)
}

func NoSetting(ctx *gin.Context) {
	data := gin.H{
		"msg":  "undefined path!",
		"code": Err,
	}
	ctx.JSON(http.StatusMethodNotAllowed, data)
}
