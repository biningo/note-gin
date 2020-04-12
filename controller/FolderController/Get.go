package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/RedisClient"
	"note-gin/pkg/utils"
	"note-gin/service/FolderService"
	"note-gin/view"
	"note-gin/view/FolderView"
)

func GetCurrentNav(c *gin.Context) {
	nav := RedisClient.GetCurrentNav()
	nav = append(nav, "Home")
	c.JSON(200, view.OkWithData("", nav))
}

func GetSubFile(c *gin.Context) {
	page := c.Param("page")
	folder_title := c.Query("title")

	folderInfos, articleInfos, total := FolderService.GetSubFile(folder_title, utils.StrToInt(page))

	//导航
	var nav []string //如果是AccessFolder  则需要加载  如果是页码跳转 则不需要加载 前端还是保留以前的nav
	if page == "1" { //page=1才可能是其他目录
		nav = RedisClient.ChangeFolderNav(folder_title) //改变redis目录路径的缓存
		nav = append(nav, "Home")
	}

	resp := view.FileList{
		Folders:  folderInfos,
		Articles: articleInfos,
		Nav:      nav,
		Total:    total,
	}

	c.JSON(200, resp)
}

//编辑区目录的懒加载请求
func GetSelectFolder(c *gin.Context) {
	folder := models.Folder{}
	err := c.ShouldBind(&folder)
	utils.ErrReport(err)

	folders := folder.GetSubFolderNoPage()
	//这里出错了一个小细节  make指定长度切片就可以直接引用位置了 如果再append的方式加入元素则会重新创建空间
	FolderSelectList := make([]FolderView.FolderSelect, len(folders))
	for i := range folders {
		FolderSelectList[i] = FolderView.FolderSelect{
			Value: folders[i].ID,
			Label: folders[i].Title,
			Leaf:  folders[i].CountSubFolder() <= 0,
		}
	}
	c.JSON(200, view.OkWithData("", FolderSelectList))
}
