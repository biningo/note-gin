package model

import "time"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ImgURL string `json:"imgUrl"` //封面图片
	Status int    `json:"status"`
}

func (this *MyBook) Add() {
	this.UpdatedAt = time.Now()
	db.Create(this)
}

func (this *MyBook) Delete() {
	db.Delete(this, "id=?", this.ID)
}

func (this MyBook) GetAll() (books []MyBook) {
	db.Table("my_book").Order("status").Find(&books)
	return
}

func (this *MyBook) Save() {
	this.UpdatedAt = time.Now()
	db.Save(this)
}
