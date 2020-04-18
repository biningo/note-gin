package models

import (
	"time"
)

type BaseModel struct {
	ID          int64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
	DeletedTime time.Time
}

//增
func AddOne(obj interface{}) {

}
func AddMany(obj []interface{}) {

}

//删
func DeleteOne(obj interface{}) {

}
func DeleteMany(attr map[string]string) {

}

//查
func FindOne(obj interface{}) {

}
func FindByPage(obj interface{}, page int) {

}

//改
func Update(attr map[string]string) {

}
func UpdateMany(obj interface{}, attr map[string]string) {

}
