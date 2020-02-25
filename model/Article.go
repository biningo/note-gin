package model

import (
	"errors"
	"note-gin/Interface/CacheCount"
	"note-gin/utils"
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

	db.Table("article").Where("deleted=?", 0).Order("updated_at desc").Offset((page - 1) * 20).Limit(20).
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
func (this *Article) Update(RedisClient CacheCount.CacheCountImpl) {
	if this.ID != 0 {
		article := Article{}
		db.Where("id=?", this.ID).Select([]string{"folder_id"}).First(&article)

		RedisClient.SubtractSubFileTotal(article.FolderID, 1)
		RedisClient.AddSubFileTotal(this.FolderID, 1)

		db.Save(&this)

	} else {
		db.Create(this)
	}
}

//Delete
func (this Article) Delete(RedisClient CacheCount.CacheCount) {
	RedisClient.SetSubFileTotal(this.FolderID, 1)

	this.Deleted = true
	this.DeletedTime = time.Now()
	db.Where("id=?", this.ID).Assign(this).FirstOrCreate(&this)
}
func (this Article) DeleteMany(ids []string, RedisClient CacheCount.CacheCount) {
	count := 0
	db.Table("article").Where("id in (?)", ids).Delete(&this).Count(&count)
	for _, v := range ids {
		vv := utils.StrToInt(v)
		RedisClient.SubtractSubFileTotal(int64(vv), 1)
	}
}

func (this Article) DeleteForever() {
	db.Where("deleted <> 0").Delete(&Article{})
}

func (this Article) Recover(RedisClient CacheCount.CacheCount) error {

	RedisClient.AddSubFileTotal(this.FolderID, 1)

	hasFolder := 0
	db.First(&this)

	db.Table("folder").Where("id=?", this.FolderID).Count(&hasFolder)
	if hasFolder != 0 {
		db.Table("article").Where("id=?", this.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}

}
