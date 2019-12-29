package main

import "github.com/go-redis/redis"

import "fmt"

// redis连接示例

// 声明一个全局的redisdb变量 (reids连接池)
var redisdb *redis.Client

// 初始化连接
func initClient() (err error) {
	// 创建一个新的连接
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 没有密码就不填
		DB:       0,  // redis分16个库(0-15)，这一步是选择在哪个库里面做
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}

	return
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("redis连接失败, err:", err)
		return
	}
	fmt.Println("redis连接成功")
}
