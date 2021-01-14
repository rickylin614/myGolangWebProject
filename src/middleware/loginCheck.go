package middleware

import "github.com/gin-gonic/gin"

func LoginCheck(ctx *gin.Context) {
	//before

	ctx.Next()
	//after

}
