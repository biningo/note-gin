package model

import (
	"errors"

	"log"
	"time"
)

type Article struct {
	BaseModel
	Title    string `form:"title" json:"title"`
	FolderID int64  `form:"folder_id" json:"folder_id"`
	MkValue  string `form:"mkValue" json:"mkValue" type:"text"`
}

//Find
func (this Article) Count() int {
	total := 0
	db.Table("article").Where("deleted=?", 0).Count(&total)
	return total
}
func (this Article) GetMany(page int) (articles []Article) {

	db.Table("article").Order("updated_at desc").Offset((page-1)*20).Limit(20).Find(&articles, "deleted=?", 0)
	return
}

func (this *Article) GetArticleInfo() {
	db.Where("id=?", this.ID).First(&this)
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
		db.Where("id=?", this.ID).Assign(*this).Save(&this)
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
func (this Article) DeleteMany(ids []string) {
	db.Table("article").Where("id in (?)", ids).Delete(&this)
}

func (this Article) DeleteForever() {
	db.Where("deleted <> 0").Delete(&Article{})
}

func (this Article) Recover() error {
	hasFolder := 0
	db.First(&this)
	log.Println(this)
	db.Table("folder").Where("id=?", this.FolderID).Count(&hasFolder)
	if hasFolder != 0 {
		db.Table("article").Where("id=?", this.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}

}
