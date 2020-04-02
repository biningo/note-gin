package BookController

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/pkg/utils"
	"note-gin/view"
)



func AddBook(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Add()
	c.JSON(200, view.OkWithData("添加成功!", book))
}

func UpdateBook(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Save()
	c.JSON(200, view.OkWithMsg("修改成功!"))
}
