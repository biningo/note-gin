package RedisClient

import (
	"github.com/go-redis/redis"
	"note-gin/config"
)

//单例
var RedisClient *redis.Client
var redisConfig = config.Conf.RedisConfig

func SetUp() {

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		DB:       redisConfig.DataBaseNumber,
		Password: redisConfig.PassWord,
	})
	RedisClient = client
}
