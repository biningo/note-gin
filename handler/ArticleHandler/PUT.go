package ArticleHandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
)

func Add(c *gin.Context) {
	add := view.ArticleAdd{}
	err := c.BindJSON(&add)
	utils.ErrReport(err) //报告错误

	article := model.Article{}
	article.Title = add.Title

	log.Println(add)
	if add.FolderTitle != "Home" {
		article.FolderID = model.Folder{}.GetFolderByTitle(add.FolderTitle).ID
	}

	article.Add()

	c.JSON(200, view.OkWithData("文章创建成功！", article))
}
func Update(c *gin.Context) {
	newArticle := model.Article{}

	err := c.ShouldBind(&newArticle)
	utils.ErrReport(err)

	newArticle.Update()

	c.JSON(200, view.OkWithMsg("文章修改成功！"))

}
