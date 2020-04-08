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
	"strings"
)

func DeleteMany(c *gin.Context) {
	ids := c.QueryArray("items[]")
	models.Article{}.DeleteMany(ids)
	c.JSON(200, view.OkWithMsg("删除成功!"))
}

func GetArticleByPage(c *gin.Context) {
	page := utils.StrToInt(c.Param("page"))
	articleLists, total := ArticleService.GetArticleByPage(page)
	c.JSON(200, view.DataList{
		Items: articleLists,
		Total: int64(total),
	})
}

//显示文章请求
func GetArticleDetail(c *gin.Context) {
	articleDetail := ArticleService.GetArticleDetail(c.Param("id"))
	c.JSON(200, articleDetail)
}

func GetRubbishArticle(c *gin.Context) {
	articles := models.Article{}.GetDeletedArticle()
	resp := view.DataList{
		Items: articles,
		Total: int64(len(articles)),
	}
	c.JSON(200, resp)
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
	article_view := view.ArticleView{}
	err := c.ShouldBind(&article_view)
	utils.ErrReport(err)

	RedisClient.SaveTempEdit(article_view)
	c.JSON(200, view.OkWithMsg("文章暂存草稿箱,1天后失效！"))
}
func TempEditGet(c *gin.Context) {
	article_view := view.ArticleView{}
	RedisClient.GetTempEdit(&article_view)
	c.JSON(200, view.OkWithData("", article_view))
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
	articleManyView := view.ArticleManageView{}
	err := c.ShouldBind(&articleManyView)
	utils.ErrReport(err)
	article := articleManyView.ToArticle()
	articleView := view.ArticleSerialize(article)
	//目录路径回溯
	articleView.DirPath = append(articleView.DirPath, articleView.FolderID)   //先添加自己的根目录
	models.Folder{}.GetFolderPath(articleView.FolderID, &articleView.DirPath) //查找路径
	c.JSON(200, view.OkWithData("", articleView))
}
