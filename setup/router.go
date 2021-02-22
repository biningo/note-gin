package setup

/**
*@Author lyer
*@Date 2/20/21 15:22
*@Describe
**/
import (
	"github.com/gin-gonic/gin"
	"note-gin/middleware"
	"note-gin/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	PublicGroup := r.Group("")
	{
		router.InitPing(PublicGroup)
	}

	PrivateGroup := r.Group("")
	{
		router.InitArticleRouter(PrivateGroup)
		router.InitBookRouter(PrivateGroup)
		router.InitFolderRouter(PrivateGroup)
		router.InitImageRouter(PrivateGroup)
	}
	return r
}
