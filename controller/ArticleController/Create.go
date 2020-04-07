package ArticleController

import (
	"github.com/gin-gonic/gin"
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
	"strings"
)

func Add(c *gin.Context) {
	articleView := view.ArticleView{}
	err := c.BindJSON(&articleView)
	utils.ErrReport(err) //报告错误
	article := models.Article{}
	article.Title = articleView.Title
	if articleView.FolderTitle != "Home" {
		article.FolderID = models.Folder{}.GetFolderByTitle(articleView.FolderTitle).ID
	}
	article.Add() //这里调用的方法必须是指针类型
	articleView = view.ArticleSerialize(article)
	//目录路径获取
	articleView.DirPath = append(articleView.DirPath, articleView.FolderID)   //先添加自己的根目录
	models.Folder{}.GetFolderPath(articleView.FolderID, &articleView.DirPath) //查找路径
	c.JSON(200, view.OkWithData("文章创建成功！", articleView))
}

//上传md
func UploadArticle(c *gin.Context) {

	folder_title := c.GetHeader("Folder-Title")
	folder_id := models.Folder{}.GetFolderByTitle(folder_title).ID

	c.Request.ParseMultipartForm(32 << 20)
	for name, file := range c.Request.MultipartForm.File {
		fp, _ := file[0].Open()
		b := make([]byte, file[0].Size)
		fp.Read(b)
		article := models.Article{}

		names := strings.Split(name, ".")
		article.Title = names[0]
		typeName := names[1]
		if typeName != "md" {
			c.JSON(200, view.RespBean{
				Code: 412, //客户端为满足条件
				Msg:  "类型 " + typeName + " 不允许",
				Data: nil,
			})
			return
		}
		isExist := article.IsExist()
		if isExist != true {
			article.FolderID = folder_id
			article.MkValue = string(b)
			article.Add()
			c.JSON(200, view.OkWithMsg("添加成功："+name))
		} else {
			c.JSON(200, view.RespBean{
				Code: 412,
				Msg:  "文件 " + name + " 已经存在",
				Data: nil,
			})
		}
	}
}
