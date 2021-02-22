package service

import (
	"note-gin/model"
	"note-gin/pkg/utils"
	"note-gin/view/ArticleView"
)

/**
*@Author lyer
*@Date 2/22/21 15:49
*@Describe
**/

func ArticleDownLoad(ID string) (string, string) {
	article := GetArticleDetail(ID)
	return article.Title, article.MkValue
}

func GetArticleByPage(page int) ([]ArticleView.ArticleInfo, int) {
	articles := model.Article{}.GetMany(page)
	total := model.Article{}.Count()
	ArticleInfos := ArticleView.ToArticleInfos(articles)
	return ArticleInfos, total
}

func GetArticleDetail(ID string) ArticleView.ArticleDetail {
	article := model.Article{}
	article.ID = int64(utils.StrToInt(ID))
	article.GetArticleInfo()
	articleDetail := ArticleView.ToArticleDetail(article)
	return articleDetail
}

func ClearRubbish() {
	model.Article{}.ClearRubbish()
}

func Delete(ID string) int64 {
	article := model.Article{}
	article.ID = int64(utils.StrToInt(ID))
	article.Delete()
	return article.ID
}