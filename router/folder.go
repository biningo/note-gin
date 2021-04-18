package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/api/v1"
)

func InitFolderRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("/folders")
	{
		BookRouter.GET("/current", v1.GetCurrentNav)
		BookRouter.GET("/files/:page", v1.GetSubFile)
		BookRouter.GET("/list", v1.GetSubFolders) //用于编辑文章选择目录时请求
		BookRouter.PATCH("/:id", v1.UpdateFolder)
		BookRouter.POST("/add", v1.AddFolder)
		BookRouter.DELETE("/:id", v1.DeleteFolder)
	}
}
