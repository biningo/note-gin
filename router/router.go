package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/handler/middleware"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //配置跨域
	r.Use(middleware.Mail()) //监控邮件

	r.Use(gin.BasicAuth(gin.Accounts{
		"test": "55555",
	}))

	r.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})
	Router = r

	FolderRouter("folder")
	ArticleRouter("article")
	QiniuRouter("qiniu")
	TimeLineRouter("time_line")
	MyRouter("my")
	return r
}
