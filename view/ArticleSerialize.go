package view

import (
	"note-gin/model"
)

func ArticleSerialize(article model.Article) (articleView ArticleView) {

	articleView.ID = article.ID
	articleView.CreatedAt = article.CreatedAt.Format("2006-01-02")
	articleView.UpdatedAt = article.UpdatedAt.Format("2006-01-02")
	articleView.DeletedAt = article.DeletedTime.Format("2006-01-02")
	articleView.FolderID = article.FolderID
	articleView.Title = article.Title
	articleView.MkValue = article.MkValue

	return
}
