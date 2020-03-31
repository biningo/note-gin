package middleware

import (
	"github.com/robfig/cron"
	"note-gin/middleware/RedisClient"
	"note-gin/utils"
)

func SendMailFromRedis() {
	Cron := cron.New()

	//每天一次
	Cron.AddFunc("0 0 22 * * ?", func() {
		AccessRecord := RedisClient.GetAccessRecord()
		RedisClient.DelAccessRecord()
		s := "今日访问量【链接访问量】：" + AccessRecord
		utils.SendMail([]string{"1535100050@qq.com"}, "note-gin", s)
	})

	Cron.Start()
}
