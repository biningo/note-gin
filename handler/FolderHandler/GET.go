package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/utils/RedisClient"
	"note-gin/view"
)

func GetSubFile(c *gin.Context) {
	pageStr := c.Param("page")
	folder := model.Folder{}
	err := c.ShouldBindQuery(&folder)
	utils.ErrReport(err)

	//导航

	nav := RedisClient.ChangeFolderNav(folder)
	nav = append(nav, "Home")

	pageNum := utils.StrToInt(pageStr)
	if folder.ID == 0 && folder.Title != "Home" {
		folder = folder.GetFolderByTitle(folder.Title)
	}
	folders, articles := folder.GetSubFile(pageNum)

	resp := view.FileList{
		Folders:  folders,
		Articles: articles,
		Nav:      nav,
	}
	c.JSON(200, resp)
}
