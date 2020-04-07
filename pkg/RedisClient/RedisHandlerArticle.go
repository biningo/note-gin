package RedisClient

import (
	"encoding/json"
	"note-gin/pkg/utils"
	"note-gin/view"
	"time"
)

func GetTempEdit(article_view *view.ArticleView) {
	isExist := RedisClient.Exists("temp_edit").Val()
	if isExist == 1 {
		s := RedisClient.Get("temp_edit").Val()

		err := json.Unmarshal([]byte(s), article_view)
		utils.ErrReport(err)
	}
}

func SaveTempEdit(temp view.ArticleView) {
	s, _ := json.Marshal(temp)                    //直接序列化存储了 因为还需要考虑没有ID的临时编辑
	RedisClient.Set("temp_edit", s, time.Hour*24) //15天
}
func DeleteTempEdit() {
	RedisClient.Del("temp_edit")
}
