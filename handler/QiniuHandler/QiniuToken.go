package QiniuHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware/QiniuClient"
)

func QiniuToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": QiniuClient.GetToken(),
	})

}
