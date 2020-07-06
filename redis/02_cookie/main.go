package main

import (
	"Go_Learn/redis/02_cookie/model"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	_, err = redisdb.Ping().Result()
	return
}

// 检查用户是否已经登录，如果已经登录返回用户的ID
func check_token(token string) (int, error) {
	res := redisdb.HGet("login:", token)

	if res.Err() != nil {
		return 0, fmt.Errorf("none")
	}

	id, _ := res.Int()

	return id, nil
}

// 参数 用户token, 用户id, 商品id
func update_token(token string, user int, item ...int) {
	date := float64(time.Now().Unix())
	redisdb.HSet("login:", token, user) // 维持令牌和已登录用户之间的映射
	redisdb.ZAdd("recent:", &redis.Z{Score: date, Member: token}) // 记录令牌最后一次出现的时间

	if len(item) > 0 {
		key := fmt.Sprintf("viewed:%v", token)
		//// 有序集合zset用法
		//redisdb.ZAdd(key, &redis.Z{Score: date, Member: item[0]}) // 记录用户最近浏览过的商品
		//redisdb.ZRemRangeByRank(key, 0, -26) // 移除旧记录，只保留最近浏览的25个商品

		// 列表list用法
		redisdb.RPush(key, item[0]) // 记录用户最近浏览过的商品 缺点: 没有记录浏览这个商品的时候的时间
		redisdb.LTrim(key, 0, 25) // 移除就记录，只保留最近浏览的25个商品
	}
}

// 去除多余的sessions
func clean_sessions() {
	var end_index int64

	for {
		size := redisdb.ZCard("recent:").Val() // 检查现有令牌数量
		if size <= model.LIMIT {
			time.Sleep(time.Second)
			continue
		}

		if size < model.LIMIT {
			end_index = size
		} else {
			end_index = model.LIMIT
		}

		tokens := redisdb.ZRange("recent:", 0, end_index-1).Val() // 获取需要移除的令牌

		session_keys := make([]string, 0, len(tokens))
		for _,v := range tokens {
			session_keys = append(session_keys, fmt.Sprintf("viewed:%v", v))
		}

		redisdb.Del(session_keys...)
		redisdb.HDel("login:", tokens...)
		redisdb.ZRem("recent:", tokens)
	}
}

func main() {
	if err := initRedis(); err != nil {
		fmt.Println("Redis 连接失败, err:", err)
		return
	}

	fmt.Println("Redis 连接成功")

	token := "123123"

	uid, err := check_token(token)
	if err != nil {
		fmt.Println("这个用户并没有登录")
		return
	}

	update_token(token, uid, 1)

	go clean_sessions()

	time.Sleep(time.Second)
}