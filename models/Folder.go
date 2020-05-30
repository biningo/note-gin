package models

import (
	"note-gin/config"
	"time"
)

var PageSize = config.Conf.AppConfig.PageSize

type Folder struct {
	BaseModel
	Title    string
	FolderID int64 //前一个文件夹
}

//Find
func (f Folder) GetRootFolder() (roots []Folder) {
	db.Find(&roots, "folder_id=?", 0)
	return
}

//目录路径查找[root->sub]
func (f Folder) GetFolderPath(FolderID int64, DirPath *[]int64) {
	if FolderID == 0 {
		return
	}
	folder := Folder{}
	db.Where("id=?", FolderID).First(&folder)

	if folder.FolderID != 0 {
		*DirPath = append([]int64{folder.FolderID}, *DirPath...)
		f.GetFolderPath(folder.FolderID, DirPath)
	} else {
		return
	}
}

func (f Folder) GetFolderByID() {
	db.Where("id=?", f.ID).First(&f)
	return
}

func (f Folder) GetSubFile(page int) (fds []Folder, articles []Article, total int) {

	fds = f.GetSubFolderOnPage(page, PageSize)
	total = f.CountSubFile()
	fdsCount := len(fds)
	if fdsCount < PageSize && fdsCount > 0 {
		//page=page-(f.CountSubFolder()/PageSize)  page-1=0
		articles = f.GetSubArticle(PageSize-fdsCount, 0)
	} else if fdsCount == 0 {
		SubFolderCount := f.CountSubFolder()
		offset := PageSize - (SubFolderCount % PageSize)
		page = page - ((SubFolderCount / PageSize) + 1)
		articles = f.GetSubArticle(PageSize, offset+(page-1)*PageSize)
	}
	return

}

func (f Folder) GetSubFolders() (folders []Folder) {
	db.Table("folder").Where("folder_id=?", f.ID).Find(&folders)
	return
}

func (f Folder) GetSubFolderOnPage(page, PageSize int) (fds []Folder) {
	db.Limit(PageSize).Offset((page-1)*PageSize).Find(&fds, "folder_id=?", f.ID)
	return
}

func (f Folder) GetSubArticle(limit, offset int) (articles []Article) {
	db.Limit(limit).Offset(offset).Where("deleted=?", 0).Select([]string{"id", "title", "updated_at", "publish_blog"}).Find(&articles, "folder_id=?", f.ID)
	return
}

func (f Folder) GetFolderInfo() {
	db.Where(f).First(&f)
}

func (f Folder) GetFolderByTitle() {
	db.Where("title=?", f.Title).First(&f)
}

//count
func (f Folder) CountSubFile() int {
	sum := f.CountSubFolder() + f.CountSubArticle()
	return sum
}

func (f Folder) CountSubFolder() (count int) {
	db.Table("folder").Where("folder_id=?", f.ID).Count(&count)
	return
}
func (f Folder) CountSubArticle() (count int) {
	db.Model(&Article{}).Where("folder_id=? and deleted=?", f.ID, 0).Count(&count)
	return
}

//Create
func (f *Folder) Add() {
	db.Create(f)
}

//Update
func (f *Folder) Update() {
	db.Model(f).Where("id=?", f.ID).Updates(map[string]interface{}{"title": f.Title, "updated_at": time.Now()})
}

//Delete递归删除
func (f *Folder) Delete() {
	db.Delete(f)
	deleteDFS(f.ID)
}
func deleteDFS(FolderID int64) {
	db.Table("article").Where("folder_id=?", FolderID).Update("deleted", true)
	sub_folder := []Folder{}
	db.Find(&sub_folder, "folder_id=?", FolderID)
	for index := range sub_folder {
		sub_folder[index].Delete()
	}
}
