package ArticleHandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"note-gin/Interface/CacheCount"
	"note-gin/middleware/RedisClient"
	"note-gin/model"
	"note-gin/utils"
	"note-gin/view"
	"strings"
)

func DeleteMany(c *gin.Context) {
	ids := c.QueryArray("items[]")
	log.Println(ids)
	model.Article{}.DeleteMany(ids, CacheCount.CacheCountImpl{})
	c.JSON(200, view.OkWithMsg("删除成功!"))
}

func GetManyArticle(c *gin.Context) {
	pageStr := c.Param("page")
	page := utils.StrToInt(pageStr)
	articles := model.Article{}.GetMany(page)
	total := model.Article{}.Count()
	articleViews := make([]view.ArticleManageView, len(articles))

	for index, v := range articles {
		articleViews[index].ID = v.ID
		articleViews[index].Title = v.Title
		articleViews[index].UpdatedAt = v.UpdatedAt.Format("2006/1/2")
	}

	c.JSON(200, view.DataList{
		Items: articleViews,
		Total: int64(total),
	})
}

func GetArticleInfo(c *gin.Context) {
	id := c.Param("id")
	article := model.Article{}
	article.ID = int64(utils.StrToInt(id))
	article.GetArticleInfo()
	c.JSON(200, gin.H{
		"mkValue": article.MkValue,
		"title":   article.Title,
		"id":      article.ID,
	})

}

func GetRubbishArticle(c *gin.Context) {
	articles := model.Article{}.GetDeletedArticle()
	resp := view.DataList{
		Items: articles,
		Total: int64(len(articles)),
	}
	c.JSON(200, resp)
}

//垃圾箱恢复
func Recover(c *gin.Context) {
	article := model.Article{}
	err := c.ShouldBind(&article)
	utils.ErrReport(err)

	ok := article.Recover(CacheCount.CacheCountImpl{})
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
	c.JSON(200, view.OkWithMsg("文章暂存草稿箱,15天后失效！"))
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

func DownLoad(c *gin.Context) {
	article := model.Article{}
	article.ID = int64(utils.StrToInt(c.Param("id")))
	article.GetArticleInfo()
	filename := article.Title

	//文件命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, strings.NewReader(article.MkValue))
}
