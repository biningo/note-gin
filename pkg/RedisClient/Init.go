package RedisClient

import "github.com/go-redis/redis"
//单例
var RedisClient *redis.Client

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     "118.178.180.115:6379",
		DB:       1,
		Password: "biningo",

	})

	RedisClient = client
}

