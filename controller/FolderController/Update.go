package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
)

func Update(c *gin.Context) {
	newFolder := models.Folder{}
	err := c.ShouldBind(&newFolder)
	utils.ErrReport(err)
	newFolder.Update()
	c.JSON(200, view.OkWithMsg("修改成功！"))
}
