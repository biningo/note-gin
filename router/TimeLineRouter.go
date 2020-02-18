package router

import "note-gin/handler/TimeLineHandler"

func TimeLineRouter(base string) {
	r := Router.Group("/" + base)
	r.POST("/add", TimeLineHandler.Add)
	r.POST("/update", TimeLineHandler.Update)
	r.GET("/delete", TimeLineHandler.Delete)

	r.GET("/all/:page", TimeLineHandler.GetAll)
}
