package model

type Article struct {
	BaseModel
	Title    string `form:"title" json:"title"`
	FolderID int64
	Tags     []Tag  `many2many:"article_tag" form:"tags" json:"tags"`
	MkValue  string `form:"mkValue" json:"mkValue" type:"text"`
	MkHtml   string `form:"mkHtml" json:"mkHtml" type:"text"`
}

//Find
func (this Article) GetArticleTag() (tags []Tag) {
	db.Model(&this).Association("tags").Find(&tags)
	return
}

func (this Article) GetArticleInfo() {
	db.Where(this).First(&this)
}
func (this Article) GetDeletedArticle() (articles []Article) {
	db.Find(&articles, "deleted=?", true)
	return
}

//Create
func (this Article) Add() {
	db.FirstOrCreate(&this, this)
}

//Update
func (this Article) Update(newArticle Article) {
	db.Where(this).Assign(newArticle).FirstOrCreate(&this)
}

//Delete
func (this Article) Delete() {
	this.Deleted = true
	db.Model(&this).Where(this).Update("deleted", true)
}

func (this Article) DeleteForever() {
	db.Delete(&Article{}, "deleted=?", true)
}
