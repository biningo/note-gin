package middleware

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware/RedisClient"
)

func Mail() gin.HandlerFunc {
	return func(context *gin.Context) {
		RedisClient.AccessRecord()
		context.Next()
	}
}
