package router

import (
	"note-gin/config"
	"note-gin/handler/TimeLineHandler"
)

func TimeLineRouter(base string) {

	r := Router.Group("/" + base)
	r.POST("/add",config.Auth,TimeLineHandler.Add)
	r.POST("/update",config.Auth,TimeLineHandler.Update)
	r.GET("/delete",config.Auth,TimeLineHandler.Delete)
	r.GET("/all/:page", TimeLineHandler.GetAll)
}
