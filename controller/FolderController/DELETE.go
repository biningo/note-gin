package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"

	"note-gin/view"
)

func Delete(c *gin.Context) {
	folder := models.Folder{}
	err := c.ShouldBind(&folder)
	utils.ErrReport(err)
	folder.Delete()
	c.JSON(200, view.OkWithData("删除目录成功！", folder.ID))
}
