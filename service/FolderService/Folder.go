package FolderService

import (
	"note-gin/models"
	"note-gin/pkg/RedisClient"
	"note-gin/pkg/utils"
	"note-gin/view/ArticleView"
	"note-gin/view/FolderView"
)

func GetFolderPath(FolderID int64, DirPath *[]int64) {
	model.Folder{}.GetFolderPath(FolderID, DirPath)
}

func GetFolderByTitle(folder_title string) FolderView.FolderInfo {
	folderInfo := FolderView.ToFolderInfo(model.Folder{Title: folder_title})
	return folderInfo
}

func GetSubFile(folder_title string, page int) ([]FolderView.FolderInfo, []ArticleView.ArticleInfo, int) {
	folder := model.Folder{}

	folder.Title = folder_title
	folder.GetFolderByTitle()

	folders, articles, total := folder.GetSubFile(page) //根据页码查找这个目录下的全部文件 total
	articleInfos := ArticleView.ToArticleInfos(articles)
	folderInfos := FolderView.ToFolderInfos(folders)

	return folderInfos, articleInfos, total

}

func ChangeNav(page string, folder_title string) []string {
	var nav []string //如果是AccessFolder  则需要加载  如果是页码跳转 则不需要加载 前端还是保留以前的nav
	if page == "1" { //page=1才可能是其他目录
		nav = RedisClient.ChangeFolderNav(folder_title) //改变redis目录路径的缓存
		nav = append(nav, "Home")
	}
	return nav
}

func GetSubFolders(id string) []FolderView.FolderSelectView {
	folder := model.Folder{}
	folder.ID = int64(utils.StrToInt(id))
	folders := folder.GetSubFolders()
	//这里出错了一个小细节  make指定长度切片就可以直接引用位置了 如果再append的方式加入元素则会重新创建空间
	folderSelectList := make([]FolderView.FolderSelectView, len(folders))
	for i := range folders {
		folderSelectList[i] = FolderView.FolderSelectView{
			Value: folders[i].ID,
			Label: folders[i].Title,
			Leaf:  folders[i].CountSubFolder() <= 0,
		}
	}
	return folderSelectList
}

func Update(folderInfo FolderView.FolderInfo) {
	folder := FolderView.ToFolder(folderInfo)
	folder.Update()
}

func Add(title string, fatherTitle string) {
	folder := model.Folder{}
	folder.Title = title

	father := model.Folder{}
	father.Title = fatherTitle
	father.GetFolderByTitle()

	folder.FolderID = father.ID
	folder.Add()
}

func Delete(id string) int64 {
	folder := model.Folder{}
	folder.ID = int64(utils.StrToInt(id))
	folder.Delete()
	return folder.ID
}
