package router

import (
	"note-gin/config"
	"note-gin/handler/MyHandler"
)

func MyRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/book/get/all",MyHandler.GetAllBook) //不设分页 直接从redis里面取 不判断是否存在
	r.GET("/book/delete/:id",config.Auth,MyHandler.DeleteBook)
	r.POST("/book/add",config.Auth,MyHandler.AddBook)
	r.POST("/book/update",config.Auth,MyHandler.Update)

	r.GET("/mood/get/many/:page",MyHandler.GetManyMood) //
	r.GET("/mood/delete",config.Auth,MyHandler.DeleteMood)
	r.GET("/mood/delete/many",config.Auth,MyHandler.DeleteManyMood)
	r.POST("/mood/add",config.Auth,MyHandler.AddMood)

}
