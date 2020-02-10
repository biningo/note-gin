package FolderHandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func GetSubFile(c *gin.Context) {
	pageStr := c.Param("page")
	folder := model.Folder{}
	err := c.ShouldBind(&folder)
	fmt.Println(folder)
	utils.ErrReport(err)

	pageNum := utils.StrToInt(pageStr)
	folders, articles := folder.GetSubFile(pageNum)

	resp := view.FileList{
		Folders:  folders,
		Articles: articles,
	}
	c.JSON(200, resp)
}
