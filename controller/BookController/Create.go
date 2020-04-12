package BookController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/logging"
	"note-gin/pkg/utils"
	"note-gin/view"
)

func AddBook(c *gin.Context) {
	book := models.MyBook{}
	err := c.ShouldBind(&book)
	logging.Error(err.Error())
	book.Add()
	c.JSON(HttpCode.SUCCESS, view.OkWithData("添加成功!", book))
}

func UpdateBook(c *gin.Context) {
	book := models.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Save()
	c.JSON(HttpCode.SUCCESS, view.OkWithMsg("修改成功!"))
}
