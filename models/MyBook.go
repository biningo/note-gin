package models

import "time"

const (
	Reading = "a在读"
	Finish  = "b读完"
	Plan    = "c想读"
)

type MyBook struct {
	ID     int64  `form:"id" json:"id"`
	Title  string `form:"title" json:"title"`
	Writer string `form:"writer" json:"writer"`
	ImgURL string `form:"img_url" json:"img_url"` //封面图片
	Status string `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}

func (book *MyBook) Add() {
	book.UpdatedAt = time.Now()
	db.Create(book)
}

func (book *MyBook) Delete() {
	db.Delete(book, "id=?", book.ID)
}

func (book MyBook) GetAll() (books []MyBook) {
	db.Table("my_book").Order("status").Find(&books)
	return
}

func (book *MyBook) Save() {
	book.UpdatedAt = time.Now()
	db.Save(book)
}
