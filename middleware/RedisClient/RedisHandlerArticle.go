package RedisClient

import (
	"encoding/json"
	"note-gin/utils"
	"note-gin/view"
	"time"
)

func GetTempEdit(article_view *view.ArticleView) {
	client := RedisInit()
	defer client.Close()

	isExist := client.Exists("temp_edit").Val()
	if isExist == 1 {
		s := client.Get("temp_edit").Val()

		err := json.Unmarshal([]byte(s), article_view)
		utils.ErrReport(err)
	}
}

func SaveTempEdit(temp view.ArticleView) {
	client := RedisInit()
	defer client.Close()

	s, _ := json.Marshal(temp)                  //直接序列化存储了 因为还需要考虑没有ID的临时编辑
	client.Set("temp_edit", s, time.Hour*24*15) //15天
}
func DeleteTempEdit() {
	client := RedisInit()
	defer client.Close()

	client.Del("temp_edit")
}
