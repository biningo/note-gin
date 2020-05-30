package router

import "note-gin/controller/AuthController"

func AuthRouter(root string){
	r:=Router.Group("/"+root)
	r.POST("/login",AuthController.Login)
}
