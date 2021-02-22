package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/api/v1"
)

func InitFolderRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("/folder")
	{
		BookRouter.GET("/current", v1.GetCurrentNav)
		BookRouter.GET("/sub/file/:page", v1.GetSubFile)
		BookRouter.GET("/sub/folder", v1.GetSubFolders) //用于编辑文章选择目录时请求
		BookRouter.PATCH("/:id", v1.Update)
		BookRouter.POST("/add", v1.Add)
		BookRouter.DELETE("/:id", v1.Delete)
	}
}
