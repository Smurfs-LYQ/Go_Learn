package main

import "github.com/go-redis/redis"

import "fmt"

// 初始化一个全局的redisdb变量
var redisdb *redis.Client

// 创建一个初始化方法
func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}

	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("Redis连接失败, err:", err)
		return
	}
	fmt.Println("Redis连接成功")

	// 设置值示例
	err = redisdb.Set("Smurfs", 18, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("设置成功")

	// 获取值示例
	val := redisdb.Get("Smurfs").Val()
	fmt.Println(val)

	// 删除值示例
	err = redisdb.Del("Smurfs").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("删除成功")
}
