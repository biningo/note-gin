package model

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDataBase(conStr string) {
	DB, err := gorm.Open("mysql", conStr)
	if err != nil {
		panic(err)
	}

	//不加s建表
	DB.SingularTable(true)

	//日志记录
	DB.LogMode(true)

	if gin.Mode() == "release" {
		DB.LogMode(false)
	}

	db = DB
	//migration(db) //迁移  首次创建数据库需要迁移创建表

}
