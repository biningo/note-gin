package RedisClient

import "github.com/go-redis/redis"

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		DB:       1,
		Password: "",

	})

	return client
}
