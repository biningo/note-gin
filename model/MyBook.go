package model

const (
	Reading = "在读"
	Finish  = "读完"
	Plan    = "想读"
)

type MyBook struct {
	ID     int64  `form:"id" json:"id"`
	Title  string `form:"title" json:"title"`
	Writer string `form:"writer" json:"writer"`
	ImgURL string `form:"img_url" json:"img_url"` //封面图片
	Status string `form:"status" json:"status"`
	Count  int    `form:"count" json:"count"`
}

func (this *MyBook) Add() {
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
	if this.Status == Finish {
		this.Count++
	}
	db.Save(this)
}
