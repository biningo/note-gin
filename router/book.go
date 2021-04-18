package router

import (
	"github.com/gin-gonic/gin"
	v1 "note-gin/api/v1"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("/books")
	{
		BookRouter.GET("/list", v1.GetAllBook) //不设分页 直接从redis里面取 不判断是否存在
		BookRouter.DELETE("/:id", v1.DeleteBook)
		BookRouter.POST("/add", v1.AddBook)
		BookRouter.PATCH("/:id", v1.UpdateBook)
	}
}
