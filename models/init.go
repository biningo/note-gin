package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"note-gin/config"
)

var db *gorm.DB

func SetUp() {
	mySqlConfig := config.Conf.MySqlConfig
	//注意添加表情的编码 并且将mysql数据库编码设置好
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		mySqlConfig.UserName, mySqlConfig.PassWord, mySqlConfig.Addr, mySqlConfig.Port, mySqlConfig.DataBaseName)

	DB, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	//不加s建表
	DB.SingularTable(true)

	if gin.Mode() == gin.ReleaseMode {
		DB.LogMode(false)
	}else if gin.Mode()==gin.DebugMode {
		DB.LogMode(true)
	}

	if config.Conf.AppConfig.MakeMigration {
		migration(db) //迁移  首次创建数据库需要迁移创建表
	}

	db = DB
}
