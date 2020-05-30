package router

import "note-gin/controller/QiniuController"

func QiniuRouter(root string) {
	r := Router.Group("/" + root)
	r.GET("/img_delete", QiniuController.ImgDelete)
	r.POST("/img_upload", QiniuController.ImgUpload)
}
