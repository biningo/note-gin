package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"note-gin/config"
)

var db *gorm.DB
var MySqlConfig = config.Conf.MySqlConfig

func SetUp() {
	//注意添加表情的编码 并且将mysql数据库编码设置好
	connStr := fmt.Sprintf("root:55555@tcp(118.178.180.115:3306)/note?charset=utf8mb4&parseTime=true",
		MySqlConfig.UserName, MySqlConfig.Port, MySqlConfig.Addr, MySqlConfig.Port, MySqlConfig.DataBaseName)

	DB, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	//不加s建表
	DB.SingularTable(true)

	if gin.Mode() == gin.ReleaseMode {
		DB.LogMode(false)
	}

	db = DB
	//migration(db) //迁移  首次创建数据库需要迁移创建表

}
