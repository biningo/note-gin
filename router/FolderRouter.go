package router

import (
	"note-gin/controller/FolderController"
)

func FolderRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/current", FolderController.GetCurrentNav)
	r.GET("/sub_file/:page", FolderController.GetSubFile)

	r.GET("/sub_folder", FolderController.GetSelectFolder) //用于编辑文章选择目录时请求 和目录导航
	r.GET("/update", FolderController.Update)
	r.GET("/add", FolderController.Add)
	r.GET("/delete", FolderController.Delete)
}
