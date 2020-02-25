package view

import "note-gin/model"

type FileList struct {
	Folders  []model.Folder
	Articles []ArticleManageView
	Nav      []string
	Total    int
}
