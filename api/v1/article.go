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
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/logging"
	"note-gin/pkg/utils"
	"note-gin/view/ArticleView"
	"note-gin/view/common"
	"strings"
)

func Add(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	logging.Error(err.Error())
	ArticleService.Add(&articleEditView)
	c.JSON(HttpCode.SUCCESS, common.OkWithData("文章创建成功！", articleEditView))
}

//上传md
func UploadArticle(c *gin.Context) {
	c.Request.ParseMultipartForm(32 << 20)
	folder_title := c.GetHeader("Folder-Title")
	file_name := ""

	isExist, ERROR := ArticleService.UploadArticle(c.Request.MultipartForm.File, folder_title, &file_name)

	if ERROR != nil && ERROR.Error() == HttpCode.HttpMsg[HttpCode.ERROR_FILE_TYPE] {
		c.JSON(200, common.RespBean{
			Code: HttpCode.ERROR_FILE_TYPE, //客户端为满足条件
			Msg:  HttpCode.HttpMsg[HttpCode.ERROR_FILE_TYPE],
			Data: nil,
		})
		return
	}

	if isExist != true {
		c.JSON(HttpCode.SUCCESS, common.OkWithMsg("添加成功："+file_name))
	} else {
		c.JSON(HttpCode.FILE_IS_EXIST_AND_UPDATE, common.RespBean{
			Code: 412,
			Msg:  "文件 " + file_name + " 已经存在;" + ERROR.Error(), //文件已经更新的警告
			Data: nil,
		})
	}
	return
}

func Delete(c *gin.Context) {
	ID := ArticleService.Delete(c.Query("id"))
	c.JSON(HttpCode.SUCCESS, common.OkWithData("成功移动到垃圾箱 定期清除哟！", ID))
}

//清空垃圾桶
func ClearRubbish(c *gin.Context) {
	ArticleService.ClearRubbish()
	c.JSON(HttpCode.SUCCESS, common.OkWithMsg("清空成功！"))
}

func DeleteMany(c *gin.Context) {
	ArticleService.DeleteMany(c.QueryArray("items[]"))
	c.JSON(200, common.OkWithMsg("删除成功!"))
}

func GetArticleByPage(c *gin.Context) {
	page := utils.StrToInt(c.Param("page"))
	articleInfos, total := ArticleService.GetArticleByPage(page)
	c.JSON(200, common.DataList{
		Items: articleInfos,
		Total: int64(total),
	})
}

//显示文章请求
func GetArticleDetail(c *gin.Context) {
	articleDetail := ArticleService.GetArticleDetail(c.Param("id"))
	c.JSON(HttpCode.SUCCESS, articleDetail)
}

func GetRubbishArticles(c *gin.Context) {
	respDataList := ArticleService.GetRubbishArticles()
	c.JSON(HttpCode.SUCCESS, respDataList)
}

//垃圾箱恢复
func ArticleRecover(c *gin.Context) {
	err := ArticleService.ArticleRecover(c.Query("id"))
	if err != nil {
		c.JSON(HttpCode.ERROR_RECOVER, common.ErrorWithMsg(HttpCode.HttpMsg[HttpCode.ERROR_RECOVER]))
	} else {
		c.JSON(200, common.OkWithMsg("恢复成功！"))
	}
}

//编辑器临时草稿保存
func TempArticleEditSave(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	utils.ErrReport(err)
	flag := ArticleService.TempArticleEditSave(articleEditView)
	if flag {
		c.JSON(HttpCode.SUCCESS, common.OkWithMsg("文章暂存草稿箱,1天后失效！"))
	} else {
		c.JSON(HttpCode.ERROR_TEMP_SAVE, common.OkWithMsg(HttpCode.HttpMsg[HttpCode.ERROR_TEMP_SAVE]))
	}
}

func TempArticleEditGet(c *gin.Context) {

	if articleEditView, ok := ArticleService.TempArticleEditGet(); ok {
		c.JSON(200, common.OkWithData("", articleEditView))
	} else {
		c.JSON(200, common.OkWithData("获取失败", articleEditView))
	}
}

func TempArticleEditDelete(c *gin.Context) {
	flag := ArticleService.TempArticleEditDelete()
	if flag == 1 {
		c.JSON(200, common.OkWithMsg("清除成功!"))
	} else {
		c.JSON(200, common.OkWithMsg("清除失败:"+string(flag)))
	}
}

func ArticleDownLoad(c *gin.Context) {
	filename, content:= ArticleService.ArticleDownLoad(c.Param("id"))
	//文件命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, strings.NewReader(content))
}

//编辑按钮点击后请求到编辑器
func Edit(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBindUri(&articleEditView)
	utils.ErrReport(err)
	ArticleService.Edit(&articleEditView)
	c.JSON(HttpCode.SUCCESS, common.OkWithData("", articleEditView))
}

//设置blog
func SetTag(c *gin.Context) {
	articleInfo := ArticleView.ArticleInfo{}
	_ = c.ShouldBind(&articleInfo)
	ArticleService.SetTag(articleInfo)
	return
}

func Update(c *gin.Context) {
	articleEditView := ArticleView.ArticleEditView{}
	err := c.ShouldBind(&articleEditView)
	logging.Error(err.Error())
	ArticleService.Update(&articleEditView)
	c.JSON(HttpCode.SUCCESS, common.OkWithData("文章保存成功！", articleEditView))
}
