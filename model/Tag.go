package model

type Tag struct {
	BaseModel
	Title string `form:"title" json:"title"`
	Articles []Article `gorm:"many2many:article_tag" form:"articles" json:"articles"`
}


