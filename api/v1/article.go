package v1

/**
*@Author lyer
*@Date 2/22/21 15:41
*@Describe
**/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"note-gin/pkg"
	"note-gin/service"
	"note-gin/view"
	"strings"
)

//DownLoad article markdown file
func ArticleDownLoad(c *gin.Context) {
	fileName, content := service.ArticleDownLoad(c.Param("id"))
	//文件命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, strings.NewReader(content))
}

//Get articles by page
func GetArticleByPage(c *gin.Context) {
	page := pkg.StrToInt(c.Param("page"))
	articleInfos, total := service.GetArticleByPage(page)
	c.JSON(200, view.DataList{
		Items: articleInfos,
		Total: int64(total),
	})
}

//Get article content
func GetArticleDetail(c *gin.Context) {
	articleDetail := service.GetArticleDetail(c.Param("id"))
	c.JSON(200, articleDetail)
}

//Add article
func AddArticle(c *gin.Context) {
	articleEditView := view.ArticleEditView{}
	c.ShouldBind(&articleEditView)
	if err := service.AddArticle(&articleEditView); err != nil {
		c.JSON(400, view.ErrorWithMsg("文章创建失败"))
		return
	}
	c.JSON(200, view.OkWithData("文章创建成功！", articleEditView))
}

//Clear rubbish
func ClearRubbish(c *gin.Context) {
	service.ClearRubbish()
	c.JSON(200, view.OkWithMsg("清空成功！"))
}

//Delete article by id
func DeleteArticle(c *gin.Context) {
	ID := service.DeleteArticle(c.Param("id"))
	c.JSON(200, view.OkWithData("成功移动到垃圾箱 定期清除哟！", ID))
}

//Delete articles by ids
func DeleteArticles(c *gin.Context) {
	service.DeleteArticles(c.QueryArray("items[]"))
	c.JSON(200, view.OkWithMsg("删除成功!"))
}

//Upload article markdown file
func UploadArticle(c *gin.Context) {
	c.Request.ParseMultipartForm(32 << 20)
	folder_title := c.GetHeader("Folder-Title")
	file_name := ""

	isExist, ERROR := service.UploadArticle(c.Request.MultipartForm.File, folder_title, &file_name)

	if ERROR != nil && ERROR.Error() == view.HttpMsg[view.ERROR_FILE_TYPE] {
		c.JSON(200, view.RespBean{
			Code: view.ERROR_FILE_TYPE, //客户端为满足条件
			Msg:  view.HttpMsg[view.ERROR_FILE_TYPE],
			Data: nil,
		})
		return
	}

	if isExist != true {
		c.JSON(200, view.OkWithMsg("添加成功："+file_name))
	} else {
		c.JSON(view.FILE_IS_EXIST_AND_UPDATE, view.RespBean{
			Code: 412,
			Msg:  "文件 " + file_name + " 已经存在;" + ERROR.Error(), //文件已经更新的警告
			Data: nil,
		})
	}
	return
}

//Get Rubbish articles
func GetRubbishArticles(c *gin.Context) {
	respDataList := service.GetRubbishArticles()
	c.JSON(200, respDataList)
}

//Rubbish recover
func ArticleRecover(c *gin.Context) {
	err := service.ArticleRecover(c.Param("id"))
	if err != nil {
		c.JSON(view.ERROR_RECOVER, view.ErrorWithMsg(view.HttpMsg[view.ERROR_RECOVER]))
	} else {
		c.JSON(200, view.OkWithMsg("恢复成功！"))
	}
}

//Get last temp article 获取上一次编辑过的草稿文件
func GetTempArticle(c *gin.Context) {

	if articleEditView, ok := service.GetTempArticle(); ok {
		c.JSON(200, view.OkWithData("", articleEditView))
	} else {
		c.JSON(200, view.OkWithData("获取失败", articleEditView))
	}
}

//Delete last temp article
func DelTempArticle(c *gin.Context) {
	flag := service.DelTempArticle()
	if flag == 1 {
		c.JSON(200, view.OkWithMsg("清除成功!"))
	} else {
		c.JSON(200, view.OkWithMsg("清除失败:"+string(flag)))
	}
}

//Save temp article
func SaveTempArticle(c *gin.Context) {
	articleEditView := view.ArticleEditView{}
	c.ShouldBind(&articleEditView)
	flag := service.SaveTempArticle(articleEditView)
	if flag {
		c.JSON(200, view.OkWithMsg("文章暂存草稿箱,1天后失效！"))
	} else {
		c.JSON(view.ERROR_TEMP_SAVE, view.OkWithMsg(view.HttpMsg[view.ERROR_TEMP_SAVE]))
	}
}

//Go to edit article 点击编辑请求 这里需要恢复article到编辑器
func ArticleEdit(c *gin.Context) {
	articleEditView := view.ArticleEditView{}
	c.ShouldBindUri(&articleEditView)
	service.ArticleEdit(&articleEditView)
	c.JSON(200, view.OkWithData("", articleEditView))
}

//Update article
func UpdateArticle(c *gin.Context) {
	articleEditView := view.ArticleEditView{}
	c.ShouldBind(&articleEditView)
	service.UpdateArticle(&articleEditView)
	c.JSON(200, view.OkWithData("文章保存成功！", articleEditView))
}

//set tag article的tag变化之后更新
func SetTag(c *gin.Context) {
	articleInfo := view.ArticleInfo{}
	c.ShouldBind(&articleInfo)
	service.SetTag(articleInfo)
	return
}
