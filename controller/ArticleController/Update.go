package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
	"time"
)

//设置blog
func SetTag(c *gin.Context) {
	article := models.Article{}
	_ = c.ShouldBind(&article)
	//article.PublishBlog ,_= c.GetPostForm("publish_blog")

	article.SetTag()
	return
}

func Update(c *gin.Context) {
	articleView := view.ArticleView{}
	err := c.ShouldBind(&articleView)
	utils.ErrReport(err)
	article := models.Article{}
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
