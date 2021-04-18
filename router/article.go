package router

import (
	"github.com/gin-gonic/gin"
	v1 "note-gin/api/v1"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("/articles")
	{
		ArticleRouter.GET("/download/:id", v1.ArticleDownLoad)
		ArticleRouter.GET("/list/:page", v1.GetArticleByPage)
		ArticleRouter.GET("/detail/:id", v1.GetArticleDetail)
		ArticleRouter.POST("/add", v1.AddArticle)
		ArticleRouter.GET("/clear", v1.ClearRubbish)
		ArticleRouter.DELETE("/:id", v1.DeleteArticle)
		ArticleRouter.DELETE("/list", v1.DeleteArticles)
		ArticleRouter.POST("/upload", v1.UploadArticle)
		ArticleRouter.GET("/rubbish", v1.GetRubbishArticles)
		ArticleRouter.GET("/recover/:id", v1.ArticleRecover)
		ArticleRouter.GET("/temp", v1.GetTempArticle)
		ArticleRouter.DELETE("/temp", v1.DelTempArticle)
		ArticleRouter.POST("/save", v1.SaveTempArticle)
		ArticleRouter.GET("/edit/:id", v1.ArticleEdit)
		ArticleRouter.PATCH("/update", v1.UpdateArticle)
		ArticleRouter.POST("/tag", v1.SetTag)
	}
}
