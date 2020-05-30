package AuthController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"note-gin/models"
	jwt_auth "note-gin/pkg/jwt-auth"
	"note-gin/pkg/logging"
	"note-gin/view/common"
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
		claims["login_name"] = account.LoginName
		claims["iat"] = time.Now()
		tokenStr, err := jwt_auth.CreateToken(claims)
		if err != nil {
			logging.Error(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"token": tokenStr,
		})

	}
	c.JSON(http.StatusUnauthorized, common.ErrorWithMsg("Account Error!"))
}
