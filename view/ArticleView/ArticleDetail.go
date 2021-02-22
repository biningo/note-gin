package ArticleView

import "note-gin/models"

type ArticleDetail struct {
	ID      int64  `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	MkValue string `form:"mkValue" json:"mkValue"`
}

func ToArticleDetail(article model.Article) ArticleDetail {
	articleDetail := ArticleDetail{
		ID:      article.ID,
		Title:   article.Title,
		MkValue: article.MkValue,
	}
	return articleDetail
}
