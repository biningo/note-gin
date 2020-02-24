package router

import "note-gin/handler/MyHandler"

func MyRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/book/get/all", MyHandler.GetAllBook) //不设分页
	r.GET("/book/delete/:id", MyHandler.DeleteBook)
	r.POST("/book/add", MyHandler.AddBook)
	r.POST("/book/update", MyHandler.Update)

	r.GET("/mood/get/many/:page", MyHandler.GetManyMood) //
	r.GET("/mood/delete", MyHandler.DeleteMood)
	r.GET("/mood/delete/many", MyHandler.DeleteManyMood)
	r.POST("/mood/add", MyHandler.AddMood)

}
