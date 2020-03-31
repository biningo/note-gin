package MyHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware/RedisClient"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func DeleteMood(c *gin.Context) {
	mood := model.MyMood{}
	err := c.ShouldBind(&mood)
	utils.ErrReport(err)
	c.JSON(200, view.OkWithMsg("删除成功"))
}
func DeleteManyMood(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	model.MyMood{}.DeleteMany(ids)
	c.JSON(200, view.OkWithMsg("删除成功"))
}

func DeleteBook(c *gin.Context) {
	book := model.MyBook{}
	book.ID = int64(utils.StrToInt(c.Param("id")))
	book.Delete()
	RedisClient.DeleteBook(int(book.ID))
	c.JSON(200, view.OkWithMsg("删除成功!"))
}
