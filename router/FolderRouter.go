package router

import (
	"note-gin/handler/FolderHandler"
)

func FolderRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/sub_file/:page", FolderHandler.GetSubFile)

	r.GET("/update", FolderHandler.Update)
	r.GET("/add", FolderHandler.Add)
	r.GET("/delete", FolderHandler.Delete)
}
