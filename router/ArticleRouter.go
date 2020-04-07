package router

import "note-gin/controller/ArticleController"

func ArticleRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/download/:id", ArticleController.ArticleDownLoad)
	r.GET("/many/:page", ArticleController.GetArticleByPage)
	r.GET("/get/:id", ArticleController.GetArticleInfo)
	r.GET("/delete/forever", ArticleController.DeleteForever)
	r.GET("/delete", ArticleController.Delete)
	r.GET("/delete/many", ArticleController.DeleteMany)
	r.GET("/rubbish", ArticleController.GetRubbishArticle)
	r.GET("/recover", ArticleController.Recover)
	r.GET("/temp_get", ArticleController.TempEditGet) //获取上次的编辑器保存
	r.GET("/temp_delete", ArticleController.TempEditDelete)

	r.POST("/temp_save", ArticleController.TempEditSave) //编辑器保存
	r.POST("/add", ArticleController.Add)
	r.POST("/update", ArticleController.Update)
	r.POST("/edit", ArticleController.Edit)
	r.POST("/set_publish_blog", ArticleController.SetTag)
	r.POST("/upload_md", ArticleController.UploadArticle)
}
