package router

import "note-gin/handler/ArticleHandler"

func ArticleRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/delete", ArticleHandler.Delete)

	r.POST("/add", ArticleHandler.Add)
	r.POST("/update", ArticleHandler.Update)

}
