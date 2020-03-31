package router

import (
	"note-gin/config"
	"note-gin/handler/FolderHandler"
)

func FolderRouter(base string) {
	r := Router.Group("/" + base)
	r.Use(config.Auth)
	r.GET("/current", FolderHandler.GetCurrentNav)
	r.GET("/sub_folder", FolderHandler.GetSelectFolder) //用于编辑文章选择目录时请求 和目录导航
	r.GET("/sub_file/:page", FolderHandler.GetSubFile)

	r.GET("/update", FolderHandler.Update)
	r.GET("/add", FolderHandler.Add)
	r.GET("/delete", FolderHandler.Delete)
}
