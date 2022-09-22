package config

import "github.com/gin-gonic/gin"

func CustomResponse(ctx *gin.Context, code int, message string, c any) {
	ctx.JSON(code, gin.H{
		"message": message,
		"code":    code,
		"data":    c,
	})
}
