package main

import (
	"note-gin/handler/middleware"
	"note-gin/model"
	"note-gin/router"
)

func main() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	model.InitDataBase("")
	r := router.NewRouter()
	r.Use(middleware.Cors())       //配置跨域
	r.Use(middleware.Mail())       //监控邮件
	middleware.SendMailFromRedis() //定时从redis读取访问记录 并且发送到邮箱  我个人监控本站
	r.Run(":9002")
}
