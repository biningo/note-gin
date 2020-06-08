package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/controller/QiniuController"
)

func QiniuRouter(root string,middleware...gin.HandlerFunc) {
	r := Router.Group("/" + root)
	r.Use(middleware...)
	r.GET("/img_delete", QiniuController.ImgDelete)
	r.POST("/img_upload", QiniuController.ImgUpload)
}
