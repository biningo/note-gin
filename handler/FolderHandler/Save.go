package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Add(c *gin.Context) {
	folder := model.Folder{}
	err := c.ShouldBind(&folder)
	if err != nil {
		log.Println(err)
	}

	folder.Add()

	c.JSON(200, view.RespBean{
		Code: 200,
		Msg:  "目录创建成功!",
		Data: folder,
	})

}

func Update(c *gin.Context) {

	newFolder := model.Folder{}
	err := c.ShouldBind(&newFolder)
	utils.ErrReport(err)

	newFolder.Update()
}
