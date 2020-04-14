package common

import (
	"note-gin/view/ArticleView"
	"note-gin/view/FolderView"
)

type FileList struct {
	Folders  []FolderView.FolderInfo
	Articles []ArticleView.ArticleInfo
	Nav      []string
	Total    int
}
