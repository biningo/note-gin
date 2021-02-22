package FolderView

import (
	"note-gin/models"
	"time"
)

type FolderInfo struct {
	ID        int64  `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}


func ToFolder(folderInfo FolderInfo) model.Folder {
	folder:= model.Folder{}
	folder.ID  = folderInfo.ID
	folder.Title = folderInfo.Title
	folder.UpdatedAt,_ = time.Parse("2006-01-02", folderInfo.UpdatedAt)
	return folder
}

func ToFolderInfo(folder model.Folder) FolderInfo {
	folderInfo := FolderInfo{}
	folderInfo.ID = folder.ID
	folderInfo.UpdatedAt = folder.UpdatedAt.Format("2006-01-02")
	folderInfo.Title = folder.Title
	return folderInfo
}

func ToFolderInfos(folders []model.Folder) []FolderInfo {
	folderInfos := make([]FolderInfo, len(folders))
	for i := range folders {
		folderInfos[i].ID = folders[i].ID
		folderInfos[i].UpdatedAt = folders[i].UpdatedAt.Format("2006-01-02")
		folderInfos[i].Title = folders[i].Title
	}
	return folderInfos
}
