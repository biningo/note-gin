package FolderView

import "note-gin/models"

type FolderInfo struct {
	ID        int64  `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	UpdatedAt string `form:"updated_at" json:"updated_at"`
}

func ToFolderInfo(folder models.Folder) FolderInfo {
	folderInfo := FolderInfo{}
	folderInfo.ID = folder.ID
	folderInfo.UpdatedAt = folder.UpdatedAt.Format("2006-01-02")
	folderInfo.Title = folder.Title
}

func ToFolderInfos(folders []models.Folder) []FolderInfo {
	folderInfos := make([]FolderInfo, len(folders))
	for i := range folders {
		folderInfos[i].ID = folders[i].ID
		folderInfos[i].UpdatedAt = folders[i].UpdatedAt.Format("2006-01-02")
		folderInfos[i].Title = folders[i].Title
	}
	return folderInfos
}
