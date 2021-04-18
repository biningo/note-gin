package service

import (
	"errors"
	"mime/multipart"
	"note-gin/model"
	"note-gin/pkg/RedisClient"
	"note-gin/view"
	"strconv"
	"strings"
	"time"
)

/**
*@Author lyer
*@Date 2/22/21 15:49
*@Describe
**/

func ArticleDownLoad(ID string) (string, string) {
	article := GetArticleDetail(ID)
	return article.Title, article.Content
}

func GetArticleByPage(page int) ([]view.ArticleInfo, int) {
	articles := model.Article{}.GetMany(page)
	total := model.Article{}.Count()
	ArticleInfos := view.ToArticleInfos(articles)
	return ArticleInfos, total
}

func GetArticleDetail(ID string) view.ArticleDetail {
	article := model.Article{}
	article.ID, _ = strconv.Atoi(ID)
	article.GetArticleInfo()
	articleDetail := view.ToArticleDetail(article)
	return articleDetail
}

func ClearRubbish() {
	model.Article{}.ClearRubbish()
}

func DeleteArticle(ID string) int {
	article := model.Article{}
	article.ID, _ = strconv.Atoi(ID)
	article.Delete()
	return article.ID
}

func DeleteArticles(IDs []string) {
	model.Article{}.DeleteMany(IDs)
}

func GetRubbishArticles() view.DataList {
	articles := model.Article{}.GetDeletedArticle()
	respDataList := view.DataList{
		Items: articles,
		Total: int64(len(articles)),
	}
	return respDataList
}

func ArticleRecover(ID string) error {
	article := model.Article{}
	article.ID, _ = strconv.Atoi(ID)
	return article.Recover()
}

func AddArticle(articleEditView *view.ArticleEditView) error {
	article := model.Article{}
	article.Title = articleEditView.Title
	if articleEditView.FolderTitle != "Home" {
		article.FolderID = GetFolderByTitle(articleEditView.FolderTitle).ID
	}
	article.Add() //这里调用的方法必须是指针类型
	articleEditView.FolderID = article.FolderID
	articleEditView.DirPath = append(articleEditView.DirPath, articleEditView.FolderID) //先添加自己的根目录
	model.Folder{}.GetFolderPath(articleEditView.FolderID, &articleEditView.DirPath)    //查找路径
	return nil
}

func UpdateArticle(articleEditView *view.ArticleEditView) {
	article := model.Article{}
	article.ID = articleEditView.ID
	article.UpdatedAt = time.Now()
	if len(articleEditView.DirPath) != 0 {
		article.FolderID = articleEditView.DirPath[len(articleEditView.DirPath)-1]
	}

	article.Content = articleEditView.Content
	article.Title = articleEditView.Title
	article.Update()

	articleEditView.UpdatedAt = article.UpdatedAt.Format("2006-01-02")
	articleEditView.CreatedAt = article.UpdatedAt.Format("2006-01-02")
	articleEditView.ID = article.ID
}

func ArticleEdit(articleEditView *view.ArticleEditView) {
	//目录路径回溯
	articleEditView.DirPath = append(articleEditView.DirPath, articleEditView.FolderID) //先添加自己的根目录
	service.GetFolderPath(articleEditView.FolderID, &articleEditView.DirPath)           //查找路径
}

func SetTag(articleInfo view.ArticleInfo) {
	article := view.ToArticle(articleInfo)
	article.SetTag()
}

func GetTempArticle() (view.ArticleEditView, bool) {
	articleEditView := view.ArticleEditView{}
	ok := RedisClient.GetTempEdit(&articleEditView)
	return articleEditView, ok
}

func DelTempArticle() int64 {
	return RedisClient.DeleteTempEdit()
}

func SaveTempArticle(articleEditView view.ArticleEditView) bool {
	flag := RedisClient.SaveTempEdit(articleEditView)
	if strings.ToLower(flag) == "ok" {
		return true
	} else {
		return false
	}
}

func UploadArticle(files map[string][]*multipart.FileHeader, folderTitle string, fileName *string) (bool, error) {
	folder_id := service.GetFolderByTitle(folderTitle).ID
	for name, file := range files {
		names := strings.Split(name, ".")
		typeName := names[1]
		if typeName != "md" {
			return false, errors.New(HttpCode.HttpMsg[HttpCode.ERROR_FILE_TYPE])
		}

		fp, _ := file[0].Open()
		b := make([]byte, file[0].Size)
		fp.Read(b)

		article := model.Article{}
		article.Title = names[0]
		*fileName = article.Title
		isExist := article.IsExist()
		if isExist != true {
			article.FolderID = folder_id
			article.Content = string(b)
			article.Add()
			return true, nil
		} else { //存在同名文件则更新 不管是否是在同一个目录下  【整个系统不允许出现同名文件】
			article.GetArticleInfoByTitle()
			article.FolderID = folder_id
			article.Content = string(b)
			article.Update()
			return false, errors.New(HttpCode.HttpMsg[HttpCode.FILE_IS_EXIST_AND_UPDATE])
		}

	}
	return false, nil
}
