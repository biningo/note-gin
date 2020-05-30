package models

import (
	"errors"
)

type Article struct {
	BaseModel
	Title    string
	FolderID int64
	MkValue  string
	Tags     string
}

//Find
func (article Article) Count() (total int) {
	db.Table("article").Where("deleted=?", 0).Count(&total)
	return total
}

func (article Article) GetMany(page int) (articles []Article) {

	db.Table("article").Where("deleted=?", 0).Order("updated_at desc").Offset((page - 1) * 10).Limit(10).
		Select([]string{"id", "updated_at", "title"}).Find(&articles)
	return
}

func (article *Article) GetArticleInfo() {
	db.Where("id=?", article.ID).First(&article)
}
func (article *Article) GetArticleInfoByTitle() {
	db.Where("title=?", article.Title).First(&article)
}

func (article Article) GetDeletedArticle() (articles []Article) {
	db.Find(&articles, "deleted=?", 1)
	return
}

//Create
func (article *Article) Add() {
	db.Create(article)
}

//Update Or Create
func (article *Article) Update() {
	if article.ID != 0 {
		db.Save(&article)
	} else {
		db.Create(article)
	}
}
func (article *Article) SetTag() {
	db.Model(&article).Update("tags", article.Tags)
}

//回收到垃圾箱
func (article *Article) Delete() {
	db.Model(&article).Update("deleted", true)
}

//真实批量删除【后台】
func (article Article) DeleteMany(ids []string) {
	db.Table("article").Where("id in (?)", ids).Delete(&article)
}

//清空垃圾箱
func (article Article) ClearRubbish() {
	db.Where("deleted <> 0").Delete(&Article{})
}

//垃圾箱恢复
func (article Article) Recover() error {
	hasFolder := 0
	db.First(&article)
	db.Table("folder").Where("id=?", article.FolderID).Count(&hasFolder)

	if hasFolder != 0 || article.FolderID == 0 {
		db.Table("article").Where("id=?", article.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}
}

func (article Article) IsExist() bool {
	c := 0
	db.Table("article").Where("title=?", article.Title).Count(&c)
	return c > 0
}
