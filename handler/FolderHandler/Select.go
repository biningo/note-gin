package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/utils"
	"note-gin/view"
)

func GetSubFile(c *gin.Context) {
	pageStr := c.Query("page")
	pageNum := utils.StrToInt(pageStr)
	folders, articles := StaticFolder.GetSubFile(pageNum)

	resp := view.FileList{
		FolderItems:  folders,
		ArticleItems: articles,
	}
	c.JSON(200, resp)
}
