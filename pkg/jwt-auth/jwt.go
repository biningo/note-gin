package jwt_auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"net/http"
	"note-gin/config"
	"time"
)

func CreateToken(claims map[string]interface{}) (tokenStr string, err error) {
	jwtClaims := jwt.MapClaims(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err = token.SignedString([]byte(config.Conf.AppConfig.JwtSecretKey))
	return
}

func ParseTokenAndValid(r *http.Request)  error {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.Conf.AppConfig.JwtSecretKey), nil //token解析出的secret必须和这里的一样才算正确
	})
	if err != nil {
		return err
	}

	if token.Valid {

		//验证过期
		claims, _ := token.Claims.(jwt.MapClaims)


		log.Println(claims["loginname"])
		log.Println(claims["exp"].(string))
		log.Println(time.Parse("2006-01-02 15:04:05", claims["exp"].(string)))

		if exp, err := time.Parse(time.RFC3339,claims["exp"].(string)); err != nil {
			return errors.New("parse Time Error")
		} else if exp.Before(time.Now()) {
			return errors.New("账户过期")
		}
		return nil
	}

	return errors.New("token错误，非法访问")

}

func DestroyToken()  {

}

func RefreshToken(){

}