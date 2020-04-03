package view

import "note-gin/models"

type FileList struct {
	Folders  []models.Folder
	Articles []ArticleManageView
	Nav      []string
	Total    int
}
