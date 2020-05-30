package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.Cors()) //配置跨域

	r.GET("/ping",middleware.JwtAuth() ,func(context *gin.Context) {
		context.Writer.WriteString("Pong")

	})
	r.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("对不起，页面找不到")
	})

	Router = r

	FolderRouter("folder")
	ArticleRouter("article")
	QiniuRouter("qiniu")
	//TimeLineRouter("time_line")
	BookRouter("my")
	AuthRouter("auth")
	return r
}
