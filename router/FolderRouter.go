package router

import (
	"note-gin/handler/FolderHandler"
)

func FolderRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/sub_file", FolderHandler.GetSubFile)

	r.POST("/update", FolderHandler.Update)
	r.POST("/add", FolderHandler.Add)
	return
}
