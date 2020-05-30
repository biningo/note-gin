package router

import (
	"note-gin/controller/BookController"
)

func BookRouter(root string) {
	r := Router.Group("/" + root)
	r.GET("/book/get/all", BookController.GetAllBook) //不设分页 直接从redis里面取 不判断是否存在
	r.GET("/book/delete/:id", BookController.DeleteBook)
	r.POST("/book/add", BookController.AddBook)
	r.POST("/book/update", BookController.UpdateBook)



}
