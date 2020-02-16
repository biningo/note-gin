package model

import "time"

type TimeLine struct {
	ID        int64     `form:"id" json:"id"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	MkValue   string    `form:"mkValue" json:"mkValue" gorm:"type:text"`
	UserName  string    `form:"username" json:"username"`
}

//不提供修改

func (this TimeLine) GetAll(page int) (timeLines []TimeLine, total int64) {
	db.Model(&TimeLine{}).Limit(10).Offset((page - 1) * 10).Order("created_at desc").Find(&timeLines)
	db.Model(&TimeLine{}).Count(&total)
	return
}

func (this *TimeLine) Add() {
	db.Create(this)
}

func (this *TimeLine) Delete() {
	db.Delete(this)
}
