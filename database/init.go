package database

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
func InitDataBase(conStr string){
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		panic(err)
	}

	//日志记录
	db.LogMode(true)

	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	DB = db
	migration(DB)


}
