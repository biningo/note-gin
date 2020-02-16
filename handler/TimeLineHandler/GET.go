package TimeLineHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func GetAll(c *gin.Context) {
	page := c.Param("page")
	result, total := model.TimeLine{}.GetAll(utils.StrToInt(page))

	c.JSON(200, view.DataList{
		Items: result,
		Total: total,
	})
}
