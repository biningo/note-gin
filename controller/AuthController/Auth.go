package AuthController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"note-gin/models"
	jwt_auth "note-gin/pkg/jwt-auth"
	"note-gin/pkg/logging"
	"time"
)

func Login(c *gin.Context) {
	account := models.Account{}
	if err := c.ShouldBind(&account); err != nil {
		logging.Error(err)
	}

	if account.LoginName == "biningo" && account.PassWord == "55555" {
		exp := time.Now().Add(time.Minute)
		claims := make(map[string]interface{})
		claims["exp"] = exp
		claims["loginname"] = account.LoginName
		claims["iat"] = time.Now().Format("2006-01-02 15:04:05")
		tokenStr, err := jwt_auth.CreateToken(claims)
		if err != nil {
			logging.Error(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenStr,
		})
		return

	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"Msg": "Account Error!",
	})
}
