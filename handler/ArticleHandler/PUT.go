package ArticleHandler

import (
	"github.com/gin-gonic/gin"
	"log"
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

	article.Add(&article) //这里调用的方法不会将article值返回出来 因为所有方法都是非指针方法  所以这里传进去了地址

	//articleView.ID=article.ID
	//articleView.FolderID = article.FolderID
	//articleView.Title=article.Title
	//articleView.CreatedAt = article.CreatedAt.Format("2006-01-02")
	//articleView.UpdatedAt = article.UpdatedAt.Format("2006-01-02")
	//articleView.DeletedAt = article.DeletedTime.Format("2006-01-02")
	articleView = utils.ArticleSerialize(article)
	//目录路径回溯
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
	article.FolderID = articleView.FolderID
	article.MkValue = articleView.MkValue
	article.Title = articleView.Title
	article.Tags = model.Tag{}.GetTagByNames(articleView.Tags)
	article.Update()

	c.JSON(200, view.OkWithMsg("文章保存成功！"))

}
func Edit(c *gin.Context) {
	article := model.Article{}
	err := c.ShouldBindJSON(&article)
	utils.ErrReport(err)
	log.Println(article)
	articleView := utils.ArticleSerialize(article)
	//目录路径回溯
	articleView.DirPath = append(articleView.DirPath, articleView.FolderID)  //先添加自己的根目录
	model.Folder{}.GetFolderPath(articleView.FolderID, &articleView.DirPath) //查找路径
	log.Println(articleView)
	c.JSON(200, view.OkWithData("", articleView))
}
