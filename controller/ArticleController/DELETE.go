package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
)

func Delete(c *gin.Context) {
	article := models.Article{}
	err := c.ShouldBind(&article)
	utils.ErrReport(err)
	article.Delete()
	c.JSON(200, view.OkWithData("成功移动到垃圾箱 定期清除哟！", article.ID))
}

//清空垃圾桶
func DeleteForever(c *gin.Context) {
	models.Article{}.DeleteForever()
	c.JSON(200, view.OkWithMsg("清空成功！"))
}
