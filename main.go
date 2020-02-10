package main

import (
	"note-gin/model"
	"note-gin/router"
)

func main() {
	model.InitDataBase("root:55555@tcp(118.178.180.115:3306)/note?charset=utf8&parseTime=true")
	r := router.NewRouter()

	r.Run(":8080")
}
