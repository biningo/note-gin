package BookController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/view/common"
)

//Book
func GetAllBook(c *gin.Context) {
	books := models.MyBook{}.GetAll()
	c.JSON(200, common.DataList{
		Items: books,
		Total: int64(len(books)),
	})
}
