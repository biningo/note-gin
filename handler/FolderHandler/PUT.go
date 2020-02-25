package FolderHandler

import (
	"github.com/gin-gonic/gin"
	"note-gin/Interface/CacheCount"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Add(c *gin.Context) {
	folder := model.Folder{}
	folder.Title = c.Query("title")
	FolderTitle := c.Query("FolderTitle")

	folder.FolderID = model.Folder{}.GetFolderByTitle(FolderTitle).ID
	folder.Add(CacheCount.CacheCountImpl{})

	c.JSON(200, view.OkWithData("目录创建成功！", folder))

}

func Update(c *gin.Context) {

	newFolder := model.Folder{}
	err := c.ShouldBind(&newFolder)
	utils.ErrReport(err)
	newFolder.Update()
	c.JSON(200, view.OkWithMsg("修改成功！"))
}
