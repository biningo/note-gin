package model

type Tag struct {
	BaseModel
	Title    string    `form:"title" json:"title"`
	Articles []Article `gorm:"many2many:article_tag" form:"articles" json:"articles"`
}

//增加或查询
func (this Tag) GetTag() {
	db.FirstOrCreate(&this, this)
}

func (this Tag) GetAllTag() (tags []Tag) {
	db.Find(&tags)
	return
}

func (this Tag) GetRelateArticle() (articles []Article) {
	db.Model(&this).Association("articles").Find(&articles)
	return
}

func (this Tag) Update(newTag Tag) {
	db.Where(this).Assign(newTag).FirstOrCreate(&this)
}
