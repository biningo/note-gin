package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})
	Router = r

	FolderRouter("folder")
	ArticleRouter("article")
	QiniuRouter("qiniu")
	TimeLineRouter("/time_line")
	MyRouter("/my")
	return r
}
