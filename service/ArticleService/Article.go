package ArticleService

import (
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view/ArticleView"
	"strings"
)

func GetArticleDetail(ID string) ArticleView.ArticleDetail {
	article := models.Article{}
	article.ID = int64(utils.StrToInt(ID))
	article.GetArticleInfo()
	articleDetail := ArticleView.ArticleDetail{
		ID:      article.ID,
		Title:   article.Title,
		MkValue: article.MkValue,
	}

	return articleDetail
}

func ArticleDownLoad(ID string) (string, string) {
	article := GetArticleDetail(ID)
	return article.Title, article.MkValue
}

func GetArticleByPage(page int) ([]ArticleView.ArticleList, int) {
	articles := models.Article{}.GetMany(page)
	total := models.Article{}.Count()
	articleList := make([]ArticleView.ArticleList, len(articles))

	for index := range articles {
		articleList[index].ID = articles[index].ID
		articleList[index].Title = articles[index].Title
		articleList[index].UpdatedAt = articles[index].UpdatedAt.Format("2006/1/2")
		articleList[index].Tags = strings.Split(articles[index].Tags, ",")
	}
	return articleList, total
}
