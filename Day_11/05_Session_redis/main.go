package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

// InitRedis 启动redis
func initRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisDB.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initRedis()
	if err != nil {
		log.Println("Redis启动失败, err:", err)
		return
	}

	r := gin.Default()

	r.Run()
}
