package router

import (
	"note-gin/config"
	"note-gin/handler/QiniuHandler"
)

func QiniuRouter(base string) {
	r := Router.Group("/" + base)
	r.Use(config.Auth)
	r.GET("/img_delete", QiniuHandler.ImgDelete)
	r.POST("/img_upload", QiniuHandler.ImgUpload)
}
