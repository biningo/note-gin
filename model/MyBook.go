package model

const (
	Reading = "在读"
	Finish  = "读完"
	Plan    = "想读"
)

type MyBook struct {
	ID      int64  `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	Author  string `form:"author" json:"author"`
	ImgURL  string `form:"img_url" json:"img_url"` //封面图片
	Process int    `form:"process" json:"process"` //读书进度
	Status  string `form:"status" json:"status"`
}
