package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/handler/middleware"
	"note-gin/model"
	"note-gin/router"
)

func main() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	model.InitDataBase("")
	r := router.NewRouter()
	middleware.SendMailFromRedis() //定时从redis读取访问记录 并且发送到邮箱  我个人监控本站
	gin.SetMode(gin.ReleaseMode)
	r.Run(":9002")
}
