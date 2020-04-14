package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/service/FolderService"
	"note-gin/view/common"
)

func Delete(c *gin.Context) {
	id:=c.Query("id")
	ID:=FolderService.Delete(id)
	c.JSON(200, common.OkWithData("删除目录成功！",ID))
}
