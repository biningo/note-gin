package ArticleView

import (
	"note-gin/models"
	"strings"
	"time"
)

type ArticleInfo struct {
	ID        int64    `json:"id" form:"id"`
	Title     string   `json:"title" form:"title"`
	UpdatedAt string   `json:"updated_at" form:"updated_at"`
	Tags      []string `json:"blogs" form:"blogs"`
}

func ToArticleInfos(articles []model.Article) []ArticleInfo {
	ArticleInfos := make([]ArticleInfo, len(articles))

	for index := range articles {
		ArticleInfos[index].ID = articles[index].ID
		ArticleInfos[index].Title = articles[index].Title
		ArticleInfos[index].UpdatedAt = articles[index].UpdatedAt.Format("2006/1/2")
		ArticleInfos[index].Tags = strings.Split(articles[index].Tags, ",")
	}
	return ArticleInfos
}

func ToArticle(articleInfo ArticleInfo) model.Article {
	article := model.Article{}
	article.ID = articleInfo.ID
	article.Title = articleInfo.Title
	article.Tags = strings.Join(articleInfo.Tags, ",")
	article.UpdatedAt, _ = time.Parse("2006-01-02", articleInfo.UpdatedAt)
	return article
}
