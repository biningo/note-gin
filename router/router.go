package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.Cors()) //配置跨域

	r.GET("/ping",func(context *gin.Context) {
		context.Writer.WriteString("Pong")

	})
	r.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("对不起，页面找不到")
	})

	Router = r
	AuthRouter("auth")

	FolderRouter("folder", middleware.JwtAuth())
	ArticleRouter("article", middleware.JwtAuth())
	QiniuRouter("qiniu", middleware.JwtAuth())
	//TimeLineRouter("time_line")
	BookRouter("my", middleware.JwtAuth())

	return r
}
