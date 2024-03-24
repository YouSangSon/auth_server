package db

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func CacheDBInit() (*redis.Client, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
