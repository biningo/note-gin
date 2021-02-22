package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/api/v1"
)

func InitImageRouter(Router *gin.RouterGroup) {
	ImageRouter := Router.Group("/img")
	{
		ImageRouter.DELETE("/:title", v1.ImgDelete)
		ImageRouter.POST("/upload", v1.ImgUpload)
	}

}
