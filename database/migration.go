package database

import (
	"github.com/jinzhu/gorm"
	"note-gin/model"
)

func migration(DB *gorm.DB) {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&model.User{}).
		AutoMigrate(&model.Folder{}).
		AutoMigrate(&model.Article{}).
		AutoMigrate(&model.Tag{})
}
