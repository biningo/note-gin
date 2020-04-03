package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/config"
	"note-gin/models"
	"note-gin/router"
)

func init() {
	config.SetUp()
	models.SetUp()
}

func main() {
	r := router.NewRouter()
	gin.SetMode(config.Conf.SystemConfig.RunMode)
	r.Run(config.Conf.SystemConfig.Host+":", config.Conf.SystemConfig.Port)
}
