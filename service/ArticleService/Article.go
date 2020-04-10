package ArticleService

import (
	"note-gin/models"
	"note-gin/pkg/utils"
	"note-gin/view"
	"note-gin/view/ArticleView"
)

func ArticleDownLoad(ID string) (string, string) {
	article := GetArticleDetail(ID)
	return article.Title, article.MkValue
}

func GetArticleByPage(page int) ([]ArticleView.ArticleInfo, int) {
	articles := models.Article{}.GetMany(page)
	total := models.Article{}.Count()
	ArticleInfos := ArticleView.ToArticleInfos(articles)
	return ArticleInfos, total
}

func GetArticleDetail(ID string) ArticleView.ArticleDetail {
	article := models.Article{}
	article.ID = int64(utils.StrToInt(ID))
	article.GetArticleInfo()
	articleDetail := ArticleView.ToArticleDetail(article)
	return articleDetail
}

func ClearRubbish() {
	models.Article{}.ClearRubbish()
}

func Delete(ID string) int64 {
	article := models.Article{}
	article.ID = int64(utils.StrToInt(ID))
	article.Delete()
	return article.ID
}

func DeleteMany(IDs []string) {
	models.Article{}.DeleteMany(IDs)
}

func GetRubbishArticles() view.DataList {
	articles := models.Article{}.GetDeletedArticle()
	respDataList := view.DataList{
		Items: articles,
		Total: int64(len(articles)),
	}
	return respDataList
}
