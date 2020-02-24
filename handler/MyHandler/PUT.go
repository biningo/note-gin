package MyHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func AddMood(c *gin.Context) {
	mood := model.MyMood{}
	err := c.ShouldBind(&mood)
	utils.ErrReport(err)
	mood.Add()
	c.JSON(200, view.OkWithData("添加成功!", mood))
}

func AddBook(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Add()
	c.JSON(200, view.OkWithData("添加成功!", book))
}

func Update(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Save()
	c.JSON(200, view.OkWithMsg("修改成功!"))
}
