package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/utils"
	"note-gin/service/ArticleService"
	"note-gin/view"
	"note-gin/view/ArticleView"
)

//设置blog
func SetTag(c *gin.Context) {
	articleInfo := ArticleView.ArticleInfo{}
	_ = c.ShouldBind(&articleInfo)
	ArticleService.SetTag(articleInfo)
	return
}

func Update(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	utils.ErrReport(err)
	ArticleService.Update(&articleEditView)
	c.JSON(200, view.OkWithData("文章保存成功！", articleEditView))
}
