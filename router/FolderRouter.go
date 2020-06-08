package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/controller/FolderController"
)

func FolderRouter(root string,middleware...gin.HandlerFunc) {
	r := Router.Group("/" + root)
	r.Use(middleware...)

	r.GET("/current", FolderController.GetCurrentNav)
	r.GET("/sub_file/:page", FolderController.GetSubFile)
	r.GET("/sub_folder", FolderController.GetSubFolders) //用于编辑文章选择目录时请求
	r.GET("/update", FolderController.Update)
	r.GET("/add", FolderController.Add)
	r.GET("/delete", FolderController.Delete)
}
