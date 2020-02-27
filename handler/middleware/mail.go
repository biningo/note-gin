package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"note-gin/middleware/RedisClient"

	"time"
)

func Mail() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		IP := context.ClientIP()
		accessTime := time.Now().Format("2006-01-02 15:04:05")
		body := fmt.Sprintf("[%s]-【%s】:%s", accessTime, IP, path)
		RedisClient.AccessRecord(body)
	}
}
