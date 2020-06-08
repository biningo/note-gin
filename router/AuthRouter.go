package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/controller/AuthController"
)

func AuthRouter(root string,middleware...gin.HandlerFunc){
	r:=Router.Group("/"+root)
	r.Use(middleware...)
	r.POST("/login",AuthController.Login)
}
