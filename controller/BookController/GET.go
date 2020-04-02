package BookController

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/view"
)

//Book
func GetAllBook(c *gin.Context) {
	books := model.MyBook{}.GetAll()
	c.JSON(200, view.DataList{
		Items: books,
		Total: int64(len(books)),
	})
}





