package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/global"
	"note-gin/setup"
)

func RunServer(r *gin.Engine) {
	if err := r.Run(global.G_CONFIG.Server.Addr); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	global.G_VP, err = setup.InitViper()
	global.G_DB, err = setup.InitMySql()
	global.G_Redis, err = setup.InitRedis()
	if err != nil {
		panic(err)
	}
	routers := setup.InitRouter()
	RunServer(routers)
}
