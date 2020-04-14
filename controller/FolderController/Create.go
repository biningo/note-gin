package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/service/FolderService"
	"note-gin/view/common"
)

func Add(c *gin.Context) {
	Title := c.Query("title")
	FatherTitle := c.Query("FatherTitle")

	FolderService.Add(Title, FatherTitle)
	c.JSON(200, common.OkWithMsg("目录创建成功！"))
}
