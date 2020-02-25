package RedisClient

import "github.com/go-redis/redis"

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "118.178.180.115:6379",
		DB:       1,
		Password: "biningo",
	})

	return client
}
