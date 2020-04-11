package ArticleController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
func ArticleRecover(c *gin.Context) {
	err := ArticleService.ArticleRecover(c.Query("id"))
	if err != nil {
		c.JSON(200, view.ErrorWithMsg(err.Error()))
	} else {
		c.JSON(200, view.OkWithMsg("恢复成功！"))
	}
}

//编辑器临时草稿保存
func TempArticleEditSave(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	utils.ErrReport(err)
	flag := ArticleService.TempArticleEditSave(articleEditView)
	if flag {
		c.JSON(200, view.OkWithMsg("文章暂存草稿箱,1天后失效！"))
	} else {
		c.JSON(200, view.OkWithMsg("文章保存失败"))
	}
}

func TempArticleEditGet(c *gin.Context) {

	if articleEditView, ok := ArticleService.TempArticleEditGet(); ok {
		c.JSON(200, view.OkWithData("", articleEditView))
	} else {
		c.JSON(200, view.OkWithData("获取失败", articleEditView))
	}
}

func TempArticleEditDelete(c *gin.Context) {
	flag := ArticleService.TempArticleEditDelete()
	if flag == 1 {
		c.JSON(200, view.OkWithMsg("清除成功!"))
	} else {
		c.JSON(200, view.OkWithMsg("清除失败:"+string(flag)))
	}
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
	ArticleService.Edit(&articleEditView)
	c.JSON(200, view.OkWithData("", articleEditView))
}
