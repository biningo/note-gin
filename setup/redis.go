package setup

import (
	"github.com/go-redis/redis"
	"note-gin/global"
)

/**
*@Author lyer
*@Date 2/22/21 15:37
*@Describe
**/

func InitRedis() (*redis.Client, error) {
	redisCfg := global.G_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
