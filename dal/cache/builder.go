package cache

import (
	"context"
	"derma/detect/config"

	"github.com/redis/go-redis/v9"
)

func RCBuilder(db int) *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS.Addr,
		Password: config.REDIS.Password,
		DB:       db,
	})

	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}

	return RedisClient
}
