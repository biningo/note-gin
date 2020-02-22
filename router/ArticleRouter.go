package router

import "note-gin/handler/ArticleHandler"

func ArticleRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/all/:page", ArticleHandler.GetAllArticle)
	r.GET("/get/:id", ArticleHandler.GetArticleInfo)
	r.GET("/delete_forever", ArticleHandler.DeleteForever)
	r.GET("/delete", ArticleHandler.Delete)
	r.GET("/delete_all", ArticleHandler.DeleteAll)
	r.GET("/rubbish", ArticleHandler.GetRubbishArticle)
	r.GET("/recover", ArticleHandler.Recover)
	r.GET("/temp_get", ArticleHandler.TempEditGet) //获取上次的编辑器保存
	r.GET("/temp_delete", ArticleHandler.TempEditDelete)

	r.POST("/temp_save", ArticleHandler.TempEditSave) //编辑器保存
	r.POST("/add", ArticleHandler.Add)
	r.POST("/update", ArticleHandler.Update)
	r.POST("/edit", ArticleHandler.Edit)
}
