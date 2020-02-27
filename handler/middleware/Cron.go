package middleware

import (
	"github.com/robfig/cron"
	"log"
	"note-gin/middleware/RedisClient"
	"note-gin/utils"
)

func SendMailFromRedis() {
	Cron := cron.New()

	//每天一次
	Cron.AddFunc("0 0 22 * * ?", func() {
		arr := RedisClient.GetAccessRecord()
		s := ""
		for _, v := range arr {
			s = s + v + "\n"
		}
		if len(arr) > 0 {
			err := utils.SendMail([]string{"1535100050@qq.com"}, "note-gin", s)
			log.Println(err)
		}
	})

	Cron.Start()
}
