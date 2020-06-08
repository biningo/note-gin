package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	jwt_auth "note-gin/pkg/jwt-auth"
)

func JwtAuth() gin.HandlerFunc{
	return func(context *gin.Context) {
		if err:=jwt_auth.ParseTokenAndValid(context.Request);err!=nil{
			context.Abort()
			context.Writer.WriteHeader(401)
			context.JSON(http.StatusUnauthorized,gin.H{
				"Code":401,
				"Msg":err.Error(),
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
