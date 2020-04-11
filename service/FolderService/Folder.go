package FolderService

import "note-gin/models"

func GetFolderPath(FolderID int64, DirPath *[]int64) {
	models.Folder{}.GetFolderPath(FolderID, DirPath)
}
