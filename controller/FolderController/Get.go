package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/RedisClient"
	"note-gin/pkg/utils"
	"note-gin/service/FolderService"
	"note-gin/view/common"
)

func GetCurrentNav(c *gin.Context) {
	nav := RedisClient.GetCurrentNav()
	nav = append(nav, "Home")
	c.JSON(200, common.OkWithData("", nav))
}

func GetSubFile(c *gin.Context) {
	page := c.Param("page")
	folder_title := c.Query("title")
	folderInfos, articleInfos, total := FolderService.GetSubFile(folder_title, utils.StrToInt(page))
	//导航
	nav := FolderService.ChangeNav(page, folder_title)
	resp := common.FileList{
		Folders:  folderInfos,
		Articles: articleInfos,
		Nav:      nav,
		Total:    total,
	}
	c.JSON(200, resp)
}

//编辑区目录的懒加载请求
func GetSubFolders(c *gin.Context) {
	id := c.Param("id")
	folderSelectList := FolderService.GetSubFolders(id)
	c.JSON(200, common.OkWithData("", folderSelectList))
}
