package models

import (
	"errors"
)

type Article struct {
	BaseModel
	Title    string `form:"title" json:"title"`
	FolderID int64  `form:"folder_id" json:"folder_id"`
	MkValue  string `form:"mkValue" json:"mkValue"`
	Tags     string `form:"tags" json:"tags"` //发布的博客平台
}

//Find
func (this Article) Count() (total int) {
	db.Table("article").Where("deleted=?", 0).Count(&total)
	return total
}

func (this Article) GetMany(page int) (articles []Article) {

	db.Table("article").Where("deleted=?", 0).Order("updated_at desc").Offset((page - 1) * 10).Limit(10).
		Select([]string{"id", "updated_at", "title"}).Find(&articles)
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
		db.Save(&this)
	} else {
		db.Create(this)
	}
}
func (this *Article) SetTag() {
	db.Model(&this).Update("publish_blog", this.PublishBlog)
}

//回收到垃圾箱
func (this *Article) Delete() {
	db.Model(&this).Update("deleted", true)
}

//真实批量删除【后台】
func (this Article) DeleteMany(ids []string) {
	db.Table("article").Where("id in (?)", ids).Delete(&this)
}

//清空垃圾箱
func (this Article) DeleteForever() {
	db.Where("deleted <> 0").Delete(&Article{})
}

//垃圾箱恢复
func (this Article) Recover() error {
	hasFolder := 0
	db.First(&this)
	db.Table("folder").Where("id=?", this.FolderID).Count(&hasFolder)

	if hasFolder != 0 || this.FolderID == 0 {
		db.Table("article").Where("id=?", this.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}
}

func (this Article) IsExist() bool {
	c := 0
	db.Table("article").Where("title=?", this.Title).Count(&c)
	return c > 0
}
