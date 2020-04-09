package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware/RedisClient"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
	"strings"
)

func GetCurrentNav(c *gin.Context) {
	nav := RedisClient.GetCurrentNav()
	nav = append(nav, "Home")
	c.JSON(200, view.OkWithData("", nav))
}

func GetSubFile(c *gin.Context) {
	pageStr := c.Param("page")
	pageNum := utils.StrToInt(pageStr)

	folder := models.Folder{}
	err := c.ShouldBindQuery(&folder)
	utils.ErrReport(err)

	//导航
	var nav []string
	if pageNum == 1 { //page=1才可能是其他目录
		nav = RedisClient.ChangeFolderNav(folder) //改变redis目录路径的缓存
		nav = append(nav, "Home")
	}

	if folder.ID == 0 && folder.Title != "Home" { //导航跳转
		folder = folder.GetFolderByTitle(folder.Title)
	}

	folders, articles, total := folder.GetSubFile(pageNum) //根据页码查找这个目录下的全部文件 total
	manyArticles := make([]view.ArticleManageView, len(articles))
	for i := range articles {
		manyArticles[i].ID = articles[i].ID
		manyArticles[i].Title = articles[i].Title
		manyArticles[i].UpdatedAt = articles[i].UpdatedAt.Format("2006-01-02")
		blogs := strings.Split(articles[i].PublishBlog, ",")
		if len(blogs) > 0 {
			manyArticles[i].Blogs = blogs
		} else {
			manyArticles[i].Blogs = nil
		}
	}
	resp := view.FileList{
		Folders:  folders,
		Articles: manyArticles,
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
	FolderSelectList := make([]view.FolderSelect, len(folders))
	for i := range folders {
		FolderSelectList[i] = view.FolderSelect{
			Value: folders[i].ID,
			Label: folders[i].Title,
			Leaf:  folders[i].CountSubFolder() <= 0,
		}
	}
	c.JSON(200, view.OkWithData("", FolderSelectList))
}
