package ArticleHandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Delete(c *gin.Context) {
	article := model.Article{}
	err := c.ShouldBind(&article)
	utils.ErrReport(err)
	log.Println(article)
	article.Delete()
	c.JSON(200, view.OkWithData("成功移动到垃圾箱 定期清除哟！", article.ID))
}

//清空垃圾桶
func ForeverDelete(c *gin.Context) {
	model.Article{}.DeleteForever()
}
