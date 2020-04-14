package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/HttpCode"
	"note-gin/service/ArticleService"
	"note-gin/view/common"
)

func Delete(c *gin.Context) {
	ID := ArticleService.Delete(c.Query("id"))
	c.JSON(HttpCode.SUCCESS, common.OkWithData("成功移动到垃圾箱 定期清除哟！", ID))
}

//清空垃圾桶
func ClearRubbish(c *gin.Context) {
	ArticleService.ClearRubbish()
	c.JSON(HttpCode.SUCCESS, common.OkWithMsg("清空成功！"))
}
