package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-weather-api/utils"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: utils.GetEnv("REDIS_ADDR", "localhost:6379"),
})

func GetCache(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func SetCache(key string, value string, ttlSeconds int) error {
	return rdb.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}
