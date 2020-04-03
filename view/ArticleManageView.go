package view

import "note-gin/models"

type ArticleManageView struct {
	ID        int64  `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	Blogs []string `json:"blogs" form:"blogs"`
}

func (this ArticleManageView) ToArticle() models.Article {
	article := models.Article{}
	article.ID = this.ID
	article.GetArticleInfo()
	return article
}
