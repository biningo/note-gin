package models

import (
	"note-gin/config"
	"time"
)

var PageSize = config.Conf.AppConfig.PageSize

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

func (this Folder) GetSubFile(page int) (fds []Folder, articles []Article, total int) {

	fds = this.GetSubFolder(page, PageSize)
	total = this.CountSubFile()
	fdsCount := len(fds)
	if fdsCount < PageSize && fdsCount > 0 {
		//page=page-(this.CountSubFolder()/PageSize)  page-1=0
		articles = this.GetSubArticle(PageSize-fdsCount, 0)
	} else if fdsCount == 0 {
		SubFolderCount := this.CountSubFolder()
		offset := PageSize - (SubFolderCount % PageSize)
		page = page - ((SubFolderCount / PageSize) + 1)
		articles = this.GetSubArticle(PageSize, offset+(page-1)*PageSize)
	}
	return

}

func (this Folder) GetSubFolderNoPage() (folders []Folder) {
	db.Where("folder_id=?", this.ID).Find(&folders)
	return
}

func (this Folder) GetSubFolder(page, PageSize int) (fds []Folder) {
	db.Limit(PageSize).Offset((page-1)*PageSize).Find(&fds, "folder_id=?", this.ID)
	return
}

func (this Folder) GetSubArticle(limit, offset int) (articles []Article) {
	db.Limit(limit).Offset(offset).Where("deleted=?", 0).Select([]string{"id", "title", "updated_at", "publish_blog"}).Find(&articles, "folder_id=?", this.ID)
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
	db.Table("folder").Where("folder_id=?", this.ID).Count(&count)
	return
}
func (this Folder) CountSubArticle() (count int) {
	db.Model(&Article{}).Where("folder_id=? and deleted=?", this.ID, 0).Count(&count)
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
func (this *Folder) Delete() {
	db.Delete(this)
	deleteDFS(this.ID)
}
func deleteDFS(FolderID int64) {
	db.Table("article").Where("folder_id=?", FolderID).Update("deleted", true)
	sub_folder := []Folder{}
	db.Find(&sub_folder, "folder_id=?", FolderID)

	for index := range sub_folder {
		sub_folder[index].Delete()
	}
}
