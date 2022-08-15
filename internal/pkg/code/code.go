package code

import "github.com/gin-gonic/gin"

func BuildReturn(context *gin.Context, state int, message string, data interface{}) {

	context.JSON(200, gin.H{
		"status":  state,
		"message": message,
		"data":    data,
	})
	return
}
