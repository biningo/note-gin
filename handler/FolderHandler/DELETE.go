package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"

	"note-gin/view"
)

func Delete(c *gin.Context) {
	folder := model.Folder{}
	err := c.ShouldBind(&folder)
	utils.ErrReport(err)
	folder.Delete()
	c.JSON(200, view.OkWithData("删除目录成功！", folder.ID))
}
