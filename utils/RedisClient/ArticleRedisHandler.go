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
	s, _ := json.Marshal(temp)
	client.Set("temp_edit", s, time.Hour*24*15) //15天
}
func DeleteTempEdit() {
	RedisInit()
	client := RedisClient
	client.Del("temp_edit")
}
