package model

import "time"

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
	db.Find(&articles, "deleted=?", 0)
	return
}

//Create
func (this Article) Add() {
	db.Create(&this)
}

//Update
func (this Article) Update() {
	db.Where("id=?", this.ID).Assign(this).FirstOrCreate(&this)
}

//Delete
func (this Article) Delete() {
	this.Deleted = true
	this.DeletedTime = time.Now()
	db.Where("id=?", this.ID).Assign(this).FirstOrCreate(&this)
}

func (this Article) DeleteForever() {
	db.Not("deleted=?", 0).Delete(&Article{})
}
