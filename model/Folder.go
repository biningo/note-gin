package model

import (
	"note-gin/config"
	"time"
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

//目录路径查找[root->sub]
func (this Folder) GetFolderPath(FolderID int64, DirPath *[]int64) {
	if FolderID == 0 {
		return
	}
	folder := Folder{}
	db.Where("id=?", FolderID).First(&folder)

	if folder.FolderID != 0 {
		*DirPath = append([]int64{folder.FolderID}, *DirPath...)
		this.GetFolderPath(folder.FolderID, DirPath)
	} else {
		return
	}
}
func (this Folder) GetFolderByID() {
	db.Where("id=?", this.ID).First(&this)
	return
}

func (this Folder) GetSubFile(page int) (fds []Folder, articles []Article) {
	fds, fdsCount := this.GetSubFolder(page, config.PageSize)
	if fdsCount < config.PageSize && fdsCount > 0 {

		//page=page-(this.CountSubFolder()/config.PageSize)  page-1=0
		articles, _ = this.GetSubArticle(config.PageSize-fdsCount, 0)

	} else if fdsCount == 0 {
		offset := config.PageSize - (this.CountSubFolder() % config.PageSize)
		page = page - ((this.CountSubFolder() / config.PageSize) + 1)
		articles, _ = this.GetSubArticle(config.PageSize, offset+(page-1)*config.PageSize)
	}
	return
}

func (this Folder) GetSubFolderNoPage() (folders []Folder) {
	db.Where("folder_id=?", this.ID).Find(&folders)
	return
}
func (this Folder) GetSubFolder(page, PageSize int) (fds []Folder, count int) {
	db.Limit(PageSize).Offset((page-1)*PageSize).Find(&fds, "folder_id=?", this.ID).Count(&count)
	return
}
func (this Folder) GetSubArticle(limit, offset int) (articles []Article, count int) {
	db.Limit(limit).Offset(offset).Where("deleted=?", 0).Find(&articles, "folder_id=?", this.ID).Find(&count)
	return
}

func (this Folder) GetFolderInfo() {
	db.Where(this).First(&this)
}
func (this Folder) GetFolderByTitle(title string) Folder {
	db.Where("title=?", title).First(&this)
	return this
}

//count
func (this Folder) CountSubFile() int {
	sum := this.CountSubFolder() + this.CountSubArticle()
	return sum
}

func (this Folder) CountSubFolder() (count int) {
	db.Model(&Folder{}).Where("folder_id=?", this.ID).Count(&count)
	return
}
func (this Folder) CountSubArticle() (count int) {
	db.Model(&Article{}).Where("folder_id=?", this.ID).Count(&count)
	return
}

//Create
func (this *Folder) Add() {
	db.Create(this)
}

//Update
func (this *Folder) Update() {
	db.Model(this).Where("id=?", this.ID).Updates(map[string]interface{}{"title": this.Title, "updated_at": time.Now()})
}

//Delete递归删除
func deleteDFS(FolderID int64) {
	db.Model(&Article{}).Where("folder_id=?", FolderID).Update("deleted", true)
	sub_folder := []Folder{}
	db.Find(&sub_folder, "folder_id=?", FolderID)
	for _, v := range sub_folder {
		v.Delete()
	}
}

func (this *Folder) Delete() {
	db.Delete(this)
	deleteDFS(this.ID)
}
