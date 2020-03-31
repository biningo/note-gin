package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/handler/middleware"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()
//中间件中如果开启了groutine 则只允许传入context的copy

	r.Use(middleware.Cors()) //配置跨域
	r.Use(middleware.Mail()) //监控邮件



	r.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})
	r.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("对不起，页面找不到")
	})


	Router = r

	FolderRouter("folder")
	ArticleRouter("article")
	QiniuRouter("qiniu")
	TimeLineRouter("time_line")
	MyRouter("my")
	return r
}
