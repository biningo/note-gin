package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/pkg/HttpCode"
	"note-gin/pkg/logging"
	"note-gin/service/ArticleService"
	"note-gin/view/ArticleView"
	"note-gin/view/common"
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
