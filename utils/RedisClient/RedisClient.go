package RedisClient

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "118.178.180.115:6379",
	})
	RedisClient = client
	return client
}
