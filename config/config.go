package config

import "github.com/gin-gonic/gin"

const (
	Admin = iota
	User
	Visitor
)

const PageSize=13

var Auth=gin.BasicAuth(gin.Accounts{
	"test": "55555",
})