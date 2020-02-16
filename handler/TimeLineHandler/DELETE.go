package TimeLineHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Delete(c *gin.Context) {
	time_line := model.TimeLine{}
	err := c.ShouldBind(time_line)
	utils.ErrReport(err)

	time_line.Delete()
	c.JSON(200, view.OkWithMsg("删除成功!"))
}
