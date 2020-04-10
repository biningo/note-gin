package ArticleController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"note-gin/models"
	"note-gin/pkg/RedisClient"
	"note-gin/pkg/utils"
	"note-gin/service/ArticleService"
	"note-gin/view"
	"note-gin/view/ArticleView"
	"strings"
)

func DeleteMany(c *gin.Context) {
	ArticleService.DeleteMany(c.QueryArray("items[]"))
	c.JSON(200, view.OkWithMsg("删除成功!"))
}

func GetArticleByPage(c *gin.Context) {
	page := utils.StrToInt(c.Param("page"))
	articleInfos, total := ArticleService.GetArticleByPage(page)
	c.JSON(200, view.DataList{
		Items: articleInfos,
		Total: int64(total),
	})
}

//显示文章请求
func GetArticleDetail(c *gin.Context) {
	articleDetail := ArticleService.GetArticleDetail(c.Param("id"))
	c.JSON(200, articleDetail)
}

func GetRubbishArticles(c *gin.Context) {
	respDataList := ArticleService.GetRubbishArticles()
	c.JSON(200, respDataList)
}

//垃圾箱恢复
func Recover(c *gin.Context) {
	article := models.Article{}
	err := c.ShouldBind(&article)
	utils.ErrReport(err)

	ok := article.Recover()
	if ok != nil {
		c.JSON(200, view.ErrorWithMsg(ok.Error()))
	} else {
		c.JSON(200, view.OkWithMsg("恢复成功！"))
	}
}

//编辑器临时草稿保存
func TempEditSave(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	utils.ErrReport(err)

	RedisClient.SaveTempEdit(articleEditView)
	c.JSON(200, view.OkWithMsg("文章暂存草稿箱,1天后失效！"))
}
func TempEditGet(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	RedisClient.GetTempEdit(&articleEditView)
	c.JSON(200, view.OkWithData("", articleEditView))
}
func TempEditDelete(c *gin.Context) {
	RedisClient.DeleteTempEdit()
	c.JSON(200, view.OkWithMsg("清除成功!"))
}

func ArticleDownLoad(c *gin.Context) {
	filename, MkValue := ArticleService.ArticleDownLoad(c.Param("id"))
	//文件命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, strings.NewReader(MkValue))
}

//编辑按钮点击后请求到编辑器
func Edit(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBindUri(&articleEditView)
	utils.ErrReport(err)
	//目录路径回溯
	articleEditView.DirPath = append(articleEditView.DirPath, articleEditView.FolderID) //先添加自己的根目录
	models.Folder{}.GetFolderPath(articleEditView.FolderID, &articleEditView.DirPath)   //查找路径
	c.JSON(200, view.OkWithData("", articleEditView))
}
