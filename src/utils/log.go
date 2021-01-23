package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetLog() (logger *log.Logger) {
	out := gin.DefaultErrorWriter
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}
	return logger
}
