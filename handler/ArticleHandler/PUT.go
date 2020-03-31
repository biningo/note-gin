package ArticleHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
	"time"
)

func Add(c *gin.Context) {
	articleView := view.ArticleView{}
	err := c.BindJSON(&articleView)
	utils.ErrReport(err) //报告错误
	article := model.Article{}
	article.Title = articleView.Title
	if articleView.FolderTitle != "Home" {
		article.FolderID = model.Folder{}.GetFolderByTitle(articleView.FolderTitle).ID
	}
	article.Add() //这里调用的方法必须是指针类型
	articleView = view.ArticleSerialize(article)
	//目录路径获取
	articleView.DirPath = append(articleView.DirPath, articleView.FolderID)  //先添加自己的根目录
	model.Folder{}.GetFolderPath(articleView.FolderID, &articleView.DirPath) //查找路径
	c.JSON(200, view.OkWithData("文章创建成功！", articleView))
}

func Update(c *gin.Context) {
	articleView := view.ArticleView{}
	err := c.ShouldBind(&articleView)
	utils.ErrReport(err)
	article := model.Article{}
	article.ID = articleView.ID
	article.UpdatedAt = time.Now()
	if len(articleView.DirPath) != 0 {
		article.FolderID = articleView.DirPath[len(articleView.DirPath)-1]
	}
	article.MkValue = articleView.MkValue
	article.Title = articleView.Title
	article.Update()

	articleView.UpdatedAt = article.UpdatedAt.Format("2006-01-02")
	articleView.CreatedAt = article.UpdatedAt.Format("2006-01-02")
	articleView.ID = article.ID
	c.JSON(200, view.OkWithData("文章保存成功！", articleView))
}

//编辑按钮点击后请求到编辑器
func Edit(c *gin.Context) {
	articleManyView := view.ArticleManageView{}
	err := c.ShouldBind(&articleManyView)
	utils.ErrReport(err)
	article := articleManyView.ToArticle()
	articleView := view.ArticleSerialize(article)
	//目录路径回溯
	articleView.DirPath = append(articleView.DirPath, articleView.FolderID)  //先添加自己的根目录
	model.Folder{}.GetFolderPath(articleView.FolderID, &articleView.DirPath) //查找路径
	c.JSON(200, view.OkWithData("", articleView))
}

//设置blog
func SetPublishBlog(c *gin.Context){
	article:=model.Article{}
	_=c.ShouldBind(&article)
	article.SetPublishBlog()
	return
}



