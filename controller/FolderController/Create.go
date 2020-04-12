package FolderController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/view"
)

func Add(c *gin.Context) {
	folder := models.Folder{}
	folder.Title = c.Query("title")
	FolderTitle := c.Query("FolderTitle")
	folder.FolderID = models.Folder{}.GetFolderByTitle(FolderTitle).ID
	folder.Add()
	c.JSON(200, view.OkWithData("目录创建成功！", folder))
}
