package BookController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
)

func DeleteBook(c *gin.Context) {
	book := models.MyBook{}
	book.ID = int64(utils.StrToInt(c.Param("id")))
	book.Delete()
	c.JSON(200, view.OkWithMsg("删除成功!"))
}
