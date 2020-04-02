package router

import "note-gin/controller/QiniuController"

func QiniuRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/img_delete", QiniuController.ImgDelete)
	r.POST("/img_upload", QiniuController.ImgUpload)
}
