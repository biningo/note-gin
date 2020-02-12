package router

import "note-gin/handler/ArticleHandler"

func ArticleRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/delete_forever", ArticleHandler.DeleteForever)
	r.GET("/delete", ArticleHandler.Delete)
	r.GET("/rubbish", ArticleHandler.GetRubbishArticle)
	r.GET("/recover", ArticleHandler.Recover)

	r.POST("/add", ArticleHandler.Add)
	r.POST("/update", ArticleHandler.Update)

}
