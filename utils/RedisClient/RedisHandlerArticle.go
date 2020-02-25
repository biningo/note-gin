package RedisClient

import (
	"encoding/json"
	"note-gin/utils"
	"note-gin/view"
	"time"
)

func GetTempEdit(article_view *view.ArticleView) {
	RedisInit()
	client := RedisClient
	s := client.Get("temp_edit").Val()
	err := json.Unmarshal([]byte(s), article_view)
	utils.ErrReport(err)
}
func SaveTempEdit(temp view.ArticleView) {
	RedisInit()
	client := RedisClient
	s, _ := json.Marshal(temp)                  //直接序列化存储了 因为还需要考虑没有ID的临时编辑
	client.Set("temp_edit", s, time.Hour*24*15) //15天
}
func DeleteTempEdit() {
	RedisInit()
	client := RedisClient
	client.Del("temp_edit")
}
