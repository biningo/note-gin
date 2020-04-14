package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/logging"
	"note-gin/service/ArticleService"
	"note-gin/view/ArticleView"
	"note-gin/view/common"
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
	logging.Error(err.Error())
	ArticleService.Update(&articleEditView)
	c.JSON(HttpCode.SUCCESS, common.OkWithData("文章保存成功！", articleEditView))
}
