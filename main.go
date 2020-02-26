package main

import (
	"note-gin/handler/middleware"
	"note-gin/model"
	"note-gin/router"
)

func main() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	model.InitDataBase("root:YouPassWord@tcp(118.178.180.115:3306)/note?charset=utf8mb4&parseTime=true")
	r := router.NewRouter()
	r.Use(middleware.Cors()) //配置跨域
	//r.Static("/dist","D:\\Go_Code\\note-gin\\static\\dist")
	//r.Static("/js","D:\\Go_Code\\note-gin\\static\\dist\\js")
	//r.Static("/css","D:\\Go_Code\\note-gin\\static\\dist\\css")
	//r.Static("/fonts","D:\\Go_Code\\note-gin\\static\\dist\\fonts")
	//r.Static("/img","D:\\Go_Code\\note-gin\\static\\dist\\img")
	//r.Static("/media","D:\\Go_Code\\note-gin\\static\\dist\\media")

	r.Run(":9002")
}
