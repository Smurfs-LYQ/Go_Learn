package db

import "github.com/go-redis/redis"

var RedisDB *redis.Client

// InitRedis 启动redis
func InitRedis() error {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisDB.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
