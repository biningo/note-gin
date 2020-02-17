package view

import "note-gin/model"

type FileList struct {
	Folders  []model.Folder
	Articles []model.Article
	Nav      []string
	Total    int
}
