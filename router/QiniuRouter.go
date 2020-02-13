package router

import "note-gin/handler/QiniuHandler"

func QiniuRouter(base string) {
	r := Router.Group("/" + base)
	r.GET("/img_delete", QiniuHandler.ImgDelete)
	r.GET("/token", QiniuHandler.QiniuToken)
	r.POST("/img_upload", QiniuHandler.ImgUpload)
}
