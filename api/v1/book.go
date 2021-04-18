package v1

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/view"
)

/**
*@Author lyer
*@Date 2/22/21 15:41
*@Describe
**/

func AddBook(c *gin.Context) {
	book := model.Book{}
	c.ShouldBind(&book)
	c.JSON(200, view.OkWithData("添加成功!", book))
}

func UpdateBook(c *gin.Context) {
	book := model.Book{}
	c.ShouldBind(&book)
	c.JSON(200, view.OkWithMsg("修改成功!"))
}

func DeleteBook(c *gin.Context) {
	c.Param("id")
	c.JSON(200, view.OkWithMsg("删除成功!"))
}

//Get all book
func GetAllBook(c *gin.Context) {
	books := []model.Book{}
	c.JSON(200, view.DataList{
		Items: books,
		Total: int64(len(books)),
	})
}
