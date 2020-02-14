package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"log"
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

func GetSubFolder(c *gin.Context) {
	folder := model.Folder{}
	err := c.ShouldBind(&folder)
	utils.ErrReport(err)

	folders := folder.GetSubFolderNoPage()
	//这里出错了一个小细节  make指定长度切片就可以直接引用位置了 如果再append的方式加入元素则会重新创建空间
	FolderSelectList := make([]view.FolderSelect, len(folders))
	for i, v := range folders {
		FolderSelectList[i] = view.FolderSelect{
			Value: v.ID,
			Label: v.Title,
			Leaf:  v.CountSubFolder() <= 0,
		}
	}
	log.Println(FolderSelectList)
	c.JSON(200, view.OkWithData("", FolderSelectList))
}