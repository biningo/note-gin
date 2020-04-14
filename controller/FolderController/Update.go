package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/logging"
	"note-gin/service/FolderService"
	"note-gin/view/FolderView"
	"note-gin/view/common"
)

func Update(c *gin.Context) {
	folderInfo:=FolderView.FolderInfo{}
	err := c.ShouldBind(&folderInfo)
	if err!=nil{
		logging.Error(err.Error())
	}
	FolderService.Update(folderInfo)
	c.JSON(HttpCode.SUCCESS, common.OkWithMsg("修改成功！"))
}
