package router

import "note-gin/controller/ArticleController"

func ArticleRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/download/:id", ArticleController.ArticleDownLoad)
	r.GET("/many/:page", ArticleController.GetArticleByPage)
	r.GET("/get/:id", ArticleController.GetArticleDetail)
	r.GET("/clear_rubbish", ArticleController.ClearRubbish)
	r.GET("/delete", ArticleController.Delete)
	r.GET("/delete/many", ArticleController.DeleteMany)
	r.GET("/rubbish", ArticleController.GetRubbishArticles)
	//-------------------------------

	r.GET("/recover", ArticleController.Recover)
	r.GET("/temp_get", ArticleController.TempEditGet) //获取上次的编辑器保存
	r.GET("/temp_delete", ArticleController.TempEditDelete)
	r.POST("/temp_save", ArticleController.TempEditSave) //编辑器保存
	r.POST("/add", ArticleController.Add)
	r.POST("/update", ArticleController.Update)
	r.GET("/edit/:id", ArticleController.Edit)
	r.POST("/set_tag", ArticleController.SetTag)
	r.POST("/upload_md", ArticleController.UploadArticle)
}
