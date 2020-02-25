package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/Interface/CacheCount"
	"note-gin/middleware/RedisClient"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func GetCurrentNav(c *gin.Context) {
	nav := RedisClient.GetCurrentNav()
	nav = append(nav, "Home")

	c.JSON(200, view.OkWithData("", nav))
}

func GetSubFile(c *gin.Context) {
	pageStr := c.Param("page")
	pageNum := utils.StrToInt(pageStr)

	folder := model.Folder{}
	err := c.ShouldBindQuery(&folder)
	utils.ErrReport(err)

	//导航
	var nav []string
	if pageNum == 1 { //page=1才可能是其他目录
		nav = RedisClient.ChangeFolderNav(folder)
		nav = append(nav, "Home")
	}
	if folder.ID == 0 && folder.Title != "Home" { //不是根目录的话就先根据要跳转的目录名select到目录
		folder = folder.GetFolderByTitle(folder.Title)
	}
	folders, articles, total := folder.GetSubFile(pageNum, CacheCount.CacheCountImpl{}) //根据页码查找这个目录下的全部文件 total redis缓存

	manyArticles := make([]view.ArticleManageView, len(articles))
	for i, v := range articles {
		manyArticles[i].ID = v.ID
		manyArticles[i].Title = v.Title
		manyArticles[i].UpdatedAt = v.UpdatedAt.Format("2006-01-02")
	}

	resp := view.FileList{
		Folders:  folders,
		Articles: manyArticles,
		Nav:      nav,
		Total:    total,
	}

	c.JSON(200, resp)
}

func GetSelectFolder(c *gin.Context) {
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
	c.JSON(200, view.OkWithData("", FolderSelectList))
}
