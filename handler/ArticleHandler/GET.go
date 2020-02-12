package ArticleHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func GetArticleInfo(c *gin.Context) {
	article := model.Article{}
	err := c.ShouldBind(&article)
	utils.ErrReport(err)

	article.GetArticleInfo()

	c.JSON(200, view.OkWithData("", article))

}

func GetRubbishArticle(c *gin.Context) {
	articles := model.Article{}.GetDeletedArticle()
	resp := view.DataList{
		Items: articles,
		Total: len(articles),
	}
	c.JSON(200, resp)
}
