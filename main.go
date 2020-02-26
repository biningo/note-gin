package main

import (
	"note-gin/handler/middleware"
	"note-gin/model"
	"note-gin/router"
)

func main() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	model.InitDataBase("root:55555@tcp(118.178.180.115:3306)/note?charset=utf8mb4&parseTime=true")
	r := router.NewRouter()
	r.Use(middleware.Cors()) //配置跨域
	r.Run(":9002")
}
