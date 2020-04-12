package FolderService

import (
	"note-gin/models"
	"note-gin/view/ArticleView"
	"note-gin/view/FolderView"
)

func GetFolderPath(FolderID int64, DirPath *[]int64) {
	models.Folder{}.GetFolderPath(FolderID, DirPath)
}

func GetFolderByTitle(folder_title string) FolderView.FolderInfo {
	folderInfo := FolderView.ToFolderInfo(models.Folder{Title: folder_title})
	return folderInfo
}

func GetSubFile(folder_title string, page int) ([]FolderView.FolderInfo, []ArticleView.ArticleInfo, int) {
	folder := models.Folder{}

	folder.Title = folder_title
	folder.GetFolderByTitle()

	folders, articles, total := folder.GetSubFile(page) //根据页码查找这个目录下的全部文件 total
	articleInfos := ArticleView.ToArticleInfos(articles)
	folderInfos := FolderView.ToFolderInfos(folders)

	return folderInfos, articleInfos, total

}
