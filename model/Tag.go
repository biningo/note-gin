package model

import "note-gin/config"

type Tag struct {
	BaseModel
	Title    string    `form:"title" json:"title"`
	Articles []Article `gorm:"many2many:article_tag" form:"articles" json:"articles"`
}

//增加或查询
func (this Tag) GetTagByNames(arr []string) (tags []Tag) {
	for _, v := range arr {
		tag := Tag{}
		db.Where("title=?", v).FirstOrCreate(&tag)
		tags = append(tags, tag)
	}
	return
}
func (this Tag) GetTag() {
	db.FirstOrCreate(&this, this)
}

func (this Tag) GetAllTag() (tags []Tag) {
	db.Find(&tags)
	return
}

func (this Tag) GetRelateArticle(page int) (articles []Article) {
	db.Model(&this).Offset((page - 1) * config.PageSize).Limit(config.PageSize).Order("created_at desc").Association("articles").Find(&articles)
	return
}

func (this Tag) Update(newTag Tag) {
	db.Where(this).Assign(newTag).FirstOrCreate(&this)
}
