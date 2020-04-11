package RedisClient

import (
	"encoding/json"
	"note-gin/pkg/utils"
	"note-gin/view/ArticleView"
	"time"
)

func GetTempEdit(article_view *ArticleView.ArticleEditView) bool {
	isExist := RedisClient.Exists("temp_edit").Val()
	if isExist == 1 {
		s := RedisClient.Get("temp_edit").Val()

		err := json.Unmarshal([]byte(s), article_view)
		utils.ErrReport(err)
		return true
	} else {
		return false
	}
}

func SaveTempEdit(temp ArticleView.ArticleEditView) string {
	s, _ := json.Marshal(temp)                                 //直接序列化存储了 因为还需要考虑没有ID的临时编辑
	return RedisClient.Set("temp_edit", s, time.Hour*24).Val() //1天
}
func DeleteTempEdit() int64 {
	return RedisClient.Del("temp_edit").Val()
}
