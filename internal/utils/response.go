package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(msg string) gin.H {
	return gin.H{
		"Success": true,
		"Error":   msg,
	}
}

func SuccessResponse(result any, msg string) gin.H {
	return gin.H{
		"Success": false,
		"Result":  result,
		"Message": msg,
	}
}
