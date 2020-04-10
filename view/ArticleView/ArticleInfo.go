package ArticleView

import (
	"note-gin/models"
	"strings"
)

type ArticleInfo struct {
	ID        int64    `json:"id" form:"id"`
	Title     string   `json:"title" form:"title"`
	UpdatedAt string   `json:"updated_at" form:"updated_at"`
	Tags      []string `json:"blogs" form:"blogs"`
}

func ToArticleInfos(articles []models.Article) []ArticleInfo {
	ArticleInfos := make([]ArticleInfo, len(articles))

	for index := range articles {
		ArticleInfos[index].ID = articles[index].ID
		ArticleInfos[index].Title = articles[index].Title
		ArticleInfos[index].UpdatedAt = articles[index].UpdatedAt.Format("2006/1/2")
		ArticleInfos[index].Tags = strings.Split(articles[index].Tags, ",")
	}
	return ArticleInfos
}
