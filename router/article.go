package router

import (
	"github.com/gin-gonic/gin"
	v1 "note-gin/api/v1"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("/article")
	{
		ArticleRouter.GET("/download/:id", v1.ArticleDownLoad)
		ArticleRouter.GET("/list/:page", v1.GetArticleByPage)
		ArticleRouter.GET("/detail/:id", v1.GetArticleDetail)
		ArticleRouter.GET("/clear", v1.ClearRubbish)
		ArticleRouter.DELETE("/:id", v1.Delete)
		ArticleRouter.DELETE("/list", v1.DeleteMany)
		ArticleRouter.GET("/rubbish", v1.GetRubbishArticles)
		ArticleRouter.GET("/recover", v1.ArticleRecover)
		ArticleRouter.GET("/temp", v1.TempArticleEditGet) //获取上次的编辑器保存
		ArticleRouter.DELETE("/temp/:id", v1.TempArticleEditDelete)
		ArticleRouter.POST("/save", v1.TempArticleEditSave) //编辑器保存
		ArticleRouter.POST("/add", v1.Add)
		ArticleRouter.PATCH("/update", v1.Update)
		ArticleRouter.GET("/edit/:id", v1.Edit)
		ArticleRouter.POST("/tag", v1.SetTag)
		ArticleRouter.POST("/upload", v1.UploadArticle)
	}
}
