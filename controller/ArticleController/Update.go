package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
	"note-gin/view/ArticleView"

	"time"
)

//设置blog
func SetTag(c *gin.Context) {
	article := models.Article{}
	_ = c.ShouldBind(&article)
	article.SetTag()
	return
}

func Update(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	utils.ErrReport(err)
	article := models.Article{}
	article.ID = articleEditView.ID
	article.UpdatedAt = time.Now()
	if len(articleEditView.DirPath) != 0 {
		article.FolderID = articleEditView.DirPath[len(articleEditView.DirPath)-1]
	}

	article.MkValue = articleEditView.MkValue
	article.Title = articleEditView.Title
	article.Update()

	articleEditView.UpdatedAt = article.UpdatedAt.Format("2006-01-02")
	articleEditView.CreatedAt = article.UpdatedAt.Format("2006-01-02")
	articleEditView.ID = article.ID
	c.JSON(200, view.OkWithData("文章保存成功！", articleEditView))
}
