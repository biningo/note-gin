package model

import (
	"note-gin/config"
)

type Folder struct {
	BaseModel
	Title    string `form:"title" json:"title"`
	FolderID int64  //前一个文件夹
}

//Find
func (this Folder) GetRootFolder() (roots []Folder) {
	db.Find(&roots, "folder_id=?", 0)
	return
}

func (this Folder) GetSubFile(page int) (fds []Folder, articles []Article) {
	fds, fdsCount := this.GetSubFolder(page, config.PageSize)
	if fdsCount < config.PageSize && fdsCount > 0 {
		//page=page-(this.CountSubFolder()/config.PageSize)  page-1=0
		db.Limit(config.PageSize - fdsCount).Offset(0).Find(&articles)
	} else if fdsCount == 0 {
		offset := config.PageSize - (this.CountSubFolder() % config.PageSize)
		page = page - ((this.CountSubFolder() / config.PageSize) + 1)
		db.Limit(config.PageSize).Offset(offset + (page-1)*config.PageSize)
	}
	return
}

func (this Folder) GetSubFolder(page, PageSize int) (fds []Folder, count int) {
	db.Limit(PageSize).Offset((page-1)*PageSize).Find(&fds, "folder_id=?", this.FolderID).Count(&count)
	return
}

func (this Folder) GetSubArticle(page, PageSize int) (articles []Article, count int) {
	db.Limit(PageSize).Offset((page-1)*PageSize).Find(&articles, "folder_id=?", this.FolderID).Find(&count)
	return
}

func (this Folder) GetFolderInfo() {
	db.First(&this)
}

//count
func (this Folder) CountSubFile() int {
	sum := this.CountSubFolder() + this.CountSubArticle()
	return sum
}

func (this Folder) CountSubFolder() (count int) {
	db.Model(&Folder{}).Where("folder_id=?", this.FolderID).Count(&count)
	return
}
func (this Folder) CountSubArticle() (count int) {
	db.Model(&Article{}).Where("folder_id=?", this.FolderID).Count(&count)
	return
}

//Create
func (this Folder) Add() {
	db.FirstOrCreate(&this, this)
}

//Update
func (this Folder) Update(newFolder Folder) {
	db.Where(this).Assign(newFolder).FirstOrCreate(&this)
}

//Delete

func deleteDFS(FolderID uint64, fds *[]Folder) {

	add_fds := []Folder{}
	add_articles := []Article{}
	db.Find(&add_fds, "folder_id=?", FolderID)
	db.Find(&add_articles, "folder_id=?", FolderID)
	for _, v := range add_articles {
		db.Delete(v)
	}

	for _, v := range add_fds {
		(*fds) = append((*fds), v)
	}

	if len(*fds) > 1 {
		id := (*fds)[0].ID
		*fds = (*fds)[1:]
		deleteDFS(id, fds)
		db.Unscoped().Delete(&(*fds)[0])
	} else if len(*fds) == 1 {
		id := (*fds)[0].ID
		deleteDFS(id, fds)
		db.Unscoped().Delete(&(*fds)[0])
	} else {
		return
	}
}

func (this Folder) Delete() {
	deleteDFS(this.ID, &[]Folder{})
}
