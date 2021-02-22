package v1

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/utils"
	"note-gin/view/common"
)

/**
*@Author lyer
*@Date 2/22/21 15:41
*@Describe
**/


func AddBook(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	logging.Error(err.Error())
	book.Add()
	c.JSON(HttpCode.SUCCESS, common.OkWithData("添加成功!", book))
}

func UpdateBook(c *gin.Context) {
	book := model.MyBook{}
	err := c.ShouldBind(&book)
	utils.ErrReport(err)
	book.Save()
	c.JSON(HttpCode.SUCCESS, common.OkWithMsg("修改成功!"))
}


func DeleteBook(c *gin.Context) {
	book := model.MyBook{}
	book.ID = int64(utils.StrToInt(c.Param("id")))
	book.Delete()
	c.JSON(200, common.OkWithMsg("删除成功!"))
}

//Book
func GetAllBook(c *gin.Context) {
	books := model.MyBook{}.GetAll()
	c.JSON(200, common.DataList{
		Items: books,
		Total: int64(len(books)),
	})
}
