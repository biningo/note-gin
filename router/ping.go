package router

import (
	"github.com/gin-gonic/gin"
)

/**
*@Author lyer
*@Date 2/22/21 15:11
*@Describe
**/

func InitPing(Router *gin.RouterGroup) {
	PingRouter := Router.Group("/ping")
	{
		PingRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "Ok")
		})
	}
}
