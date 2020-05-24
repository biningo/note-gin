package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/config"
	"note-gin/models"
	"note-gin/pkg/QiniuClient"
	"note-gin/pkg/RedisClient"
	"note-gin/pkg/logging"
	"note-gin/router"
)

func init() {
	config.SetUp()
	models.SetUp()
	RedisClient.SetUp()
	QiniuClient.SetUp()
	logging.SetUp()
}

func main() {
	r := router.NewRouter()
	gin.SetMode(config.Conf.ServerConfig.RunMode)
	r.Run(config.Conf.ServerConfig.Host+":", config.Conf.ServerConfig.Port)

}
