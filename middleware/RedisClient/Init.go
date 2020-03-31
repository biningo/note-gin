package RedisClient

import "github.com/go-redis/redis"

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		

	})

	return client
}

