package view

import "note-gin/model"

type FileList struct {
	FolderItems  []model.Folder
	ArticleItems []model.Article
}
