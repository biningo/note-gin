package v1

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/RedisClient"
	"note-gin/service/FolderService"
	"note-gin/view"
	"note-gin/view/FolderView"
	"note-gin/view/common"
)

/**
*@Author lyer
*@Date 2/22/21 15:41
*@Describe
**/

func AddFolder(c *gin.Context) {
	//Title := c.Query("title")
	//FatherTitle := c.Query("FatherTitle")

	c.JSON(200, view.OkWithMsg("目录创建成功！"))
}

func DeleteFolder(c *gin.Context) {
	c.Param("id")
	//ID:=FolderService.Delete(id)
	c.JSON(200, view.OkWithMsg(("删除目录成功！")))
}

func GetCurrentNav(c *gin.Context) {
	nav := RedisClient.GetCurrentNav()
	nav = append(nav, "Home")
	c.JSON(200, view.OkWithData("", nav))
}

func GetSubFile(c *gin.Context) {
	page := c.Param("page")
	folder_title := c.Query("title")
	folderInfos, articleInfos, total := FolderService.GetSubFile(folder_title, pkg.StrToInt(page))
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
	c.JSON(200, view.OkWithData("", folderSelectList))
}

func UpdateFolder(c *gin.Context) {
	folderInfo := FolderView.FolderInfo{}
	err := c.ShouldBind(&folderInfo)
	if err != nil {
	}
	FolderService.Update(folderInfo)
	c.JSON(HttpCode.SUCCESS, view.OkWithMsg("修改成功！"))
}
