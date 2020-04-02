package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/router"
)

func main() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	model.InitDataBase("root:1111@tcp(118.178.180.115:3306)/note?charset=utf8mb4&parseTime=true")
	r := router.NewRouter()


	gin.SetMode(gin.ReleaseMode)
	r.Run(":9002")
}
