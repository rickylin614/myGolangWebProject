package middleware

import (
	"github.com/gin-gonic/gin"
)

func Common(ctx *gin.Context) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		s := "請求發生錯誤! 請稍後重新再試"
	// 		if gin.IsDebugging() {
	// 			s = fmt.Sprintf("err : %v", err)
	// 			fmt.Printf("err : %v\n", err)
	// 		}
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"msg":  s,
	// 			"code": "error",
	// 		})
	// 	}
	// }()
	//before
	ctx.Next()
	//after
}
