package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})

	return r
}
