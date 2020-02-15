package model

import (
	"errors"
	"time"
)

type Article struct {
	BaseModel
	Title    string `form:"title" json:"title"`
	FolderID int64  `form:"folder_id" json:"folder_id"`
	Tags     []Tag  `many2many:"article_tag" form:"tags" json:"tags"`
	MkValue  string `form:"mkValue" json:"mkValue" type:"text"`
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
	db.Find(&articles, "deleted=?", 1)
	return
}

//Create
func (this *Article) Add() {
	db.Create(this)
}

//Update Or Create
func (this *Article) Update() {
	if this.ID != 0 {
		db.Where("id=?", this.ID).Assign(*this).FirstOrCreate(this)
	} else {
		db.Create(this)
	}
}

//Delete
func (this Article) Delete() {
	this.Deleted = true
	this.DeletedTime = time.Now()
	db.Where("id=?", this.ID).Assign(this).FirstOrCreate(&this)
}

func (this Article) DeleteForever() {
	db.Where("deleted <> 0").Delete(&Article{})
}
func (this Article) Recover() error {
	hasFolder := 0
	db.First(&this)
	db.Where("id=?", this.FolderID).Count(&hasFolder)
	if hasFolder != 0 {
		db.Model(&this).Where("id=?", this.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}

}
