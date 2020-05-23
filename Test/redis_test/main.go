package main

import (
	"Go_Learn/Test/redis_test/article"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	//"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

// 声明全局MySQL连接客户端
var db *sql.DB

// 声明全局redis连接客户端
var redisdb *redis.Client

// 初始化连接MySQL
func initMySQL() (err error) {
	// dsn := "root:smurfs&tcp(127.0.0.1:3306)/test"
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/test" // 用户名:密码@连接方式(IP:端口)/库名

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
	}

	return
}

// 初始化连接redis
func initRedis() (err error) {
	// 创建一个新连接
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 设置数据库地址和端口
		Password: "",               // 设置密码，没有可以为空
		DB:       0,                // 选择连接的分库
	})

	_, err = redisdb.Ping().Result()

	return
}

// 无排序无分组
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	/*
//		1. 加载模板文件
//		2. 检查redis中是否有缓存
//			2-1. 如果有缓存直接加载到页面
//				2-1-1. 跑一个goroutine去查询本地数据库检查数据redis是否有更新
//				2-1-2. 如果输入有更新进行更新
//			2-2. 如果没有缓存从MySQL中查询数据，并载入到首页
//	*/
//
//	fmt.Println("首页-接入请求")
//
//	//1. 加载模板文件
//	t, err := template.ParseFiles("./template/index.html")
//	if err != nil {
//		fmt.Println("文件加载失败, err:", err)
//		return
//	}
//
//	var model article.Article
//	var articles =make([]article.Article, 0, 10)
//
//L1:
//	//2. 检查redis中是否有缓存
//	keys := redisdb.Keys("article:*")
//	if cap(keys.Val()) > 0 {
//		for _,v := range keys.Val() {
//			val := redisdb.HGetAll(v)
//			res, _ := json.Marshal(val.Val())
//			json.Unmarshal(res, &model)
//			articles = append(articles, model)
//		}
//	}
//
//	if len(articles) > 0 {
//		//2-1. 跑一个goroutine去查询本地数据库检查数据redis是否有更新, 如果输入有更新进行更新
//		go index_redis()
//
//		//2-2. 如果有缓存直接加载到页面
//		t.Execute(w, articles)
//	} else {
//		//2-2. 如果没有缓存从MySQL中查询数据，并载入到首页
//		index_redis()
//
//		goto L1
//	}
//}

// 可以排序
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("首页-接入请求-排序")

	//加载页面
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("页面加载失败, err: ", err)
		return
	}

	// 创建一个用于保存文章信息的切片
	var model article.Article
	articles := make([]article.Article, 0, 10)
L1:
	// 检查redis是否有缓存
	keys := redisdb.Keys("article:*")
	if cap(keys.Val()) > 0 {
		// 通过排序获取所有的key
		// 通过积分排序
		res := redisdb.ZRevRange("source:", 0, -1)
		// 通过时间排序
		//res := redisdb.ZRevRange("time:", 0, -1)

		for _,v := range res.Val() {
			article_res := redisdb.HGetAll(v).Val()

			// 将字符串时间戳转换格式化
			// 1. 先将字符串转换成int64格式的
			i, _ := strconv.ParseInt(article_res["time"], 10, 64)
			// 2. 再将转换后的数字字符串转换成time格式
			tm := time.Unix(i, 0)
			// 赋值并格式化
			article_res["time"] = tm.Format("2006-01-02 15:04:05")

			article_res_json, _ := json.Marshal(article_res)
			json.Unmarshal(article_res_json, &model)
			articles = append(articles, model)
		}
	}

	if len(articles) > 0 {
		go index_redis()

		t.Execute(w, articles)
	} else {
		index_redis()

		goto L1
	}
}

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("接收请求")
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("文件加载失败, err:", err)
		return
	}

	rows, err := db.Query("select title,id,writer,time,votes from articles")
	if err != nil {
		fmt.Println("查询失败, err:", err)
		return
	}

	defer rows.Close()

	var scrap article.Article
	var lists = make([]article.Article, 0, 10)

	for rows.Next() {
		rows.Scan(&scrap.Title, &scrap.Url, &scrap.Poster, &scrap.Time, &scrap.Votes)
		lists = append(lists, scrap)
	}

	var time_list = make([]*redis.Z, 0, 10)
	var vote_list = make([]*redis.Z, 0, 10)

	for k, v := range lists {
		date, _ := strconv.Atoi(v.Time)
		lists[k].Time = time.Unix(int64(date), 0).Format("2006-01-02 15:04:05")
		name := fmt.Sprintf("article:%s", v.Url)
		redisdb.HSet(name, "title", v.Title)
		redisdb.HSet(name, "link", v.Url)
		redisdb.HSet(name, "poster", v.Poster)
		redisdb.HSet(name, "time", v.Time)
		redisdb.HSet(name, "votes", v.Votes)

		time_list = append(time_list, &redis.Z{Score: float64(date), Member: name})
		vote_list = append(vote_list, &redis.Z{Score: float64(v.Votes), Member: name})
	}

	redisdb.ZAdd("time:", time_list...)
	redisdb.ZAdd("source:", vote_list...)

	t.Execute(w, lists)
}
*/

// 检测首页文章是否有更新
func index_redis() {
	fmt.Println("开始更新-首页文章")
	// 查询数据库中的信息，判断新添加的文章有没有添加到redis中
	res, err := db.Query("select * from articles")
	if err != nil {
		fmt.Println("更新失败-查询失败, err:", err)
		return
	}

	defer res.Close()

	var article article.Article
	var time_list = make([]*redis.Z, 0, 10)
	var vote_list = make([]*redis.Z, 0, 10)

	for res.Next() {
		err := res.Scan(&article.Url, &article.Title, &article.Poster, &article.Votes, &article.Time)
		if err != nil {
			fmt.Println("首页更新-解析失败, err:", err)
			continue
		}

		name := fmt.Sprintf("article:%v", article.Url)

		// 检测指定文章是否存在
		res := redisdb.Keys(name)
		if res.Err() == nil {
			redisdb.HSet(name, "title", article.Title)
			redisdb.HSet(name, "link", article.Url)
			redisdb.HSet(name, "poster", article.Poster)
			redisdb.HSet(name, "time", article.Time)
			redisdb.HSet(name, "votes", article.Votes)

			date,_ := strconv.Atoi(article.Time)
			vote,_ := strconv.Atoi(article.Votes)

			time_list = append(time_list, &redis.Z{float64(date), name})
			vote_list = append(vote_list, &redis.Z{float64(vote), name})
		} else {
			//更新点赞信息
			vote,_ := strconv.Atoi(article.Votes)
			redisdb.HSet(name, "votes", vote)
			redisdb.ZAdd("source:", &redis.Z{float64(vote), name})
		}
	}

	//添加新值到time:表中
	redisdb.ZAdd("time:", time_list...)
	//添加新值到source:表中
	redisdb.ZAdd("source:", vote_list...)
}

//点赞
func scrapHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("点赞请求接入")
	r.ParseForm()

	article_id := r.Form.Get("ID")
	uid := r.Form.Get("uid")
	user_id := fmt.Sprintf("user:%v", uid)
	voted := fmt.Sprintf("voted:%v", article_id)
	member := fmt.Sprintf("article:%v", article_id)

	// 检测用户是否已在已投票用户表中
	user_res := redisdb.SIsMember(voted, user_id).Val()

	if user_res {
		// 已经投过票了
		fmt.Printf("用户ID为: %v 的用户，已经为ID为: %v 投过票了\n", uid, article_id)
		http.Redirect(w, r, "/", 302)
		return
	}

	// 将投票的用户添加到已投票的集合中
	redisdb.SAdd(voted, user_id)
	// 对应增加统计文章票数的有序集合中的数值 两个表
	//积分表
	redisdb.ZIncrBy("source:", float64(1), member)
	//文章表
	redisdb.HIncrBy(member, "votes", 1)

	// 修改数据库(结合事务管理)
	tx, err := db.Begin() // 事务开始
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Println("事务开始失败, err:", err)
		http.Redirect(w, r, "/", 302)
		return
	}

	//执行修改数据库的操作
	db_res, err := db.Exec("update articles set votes = votes+1 where id = ?", article_id)
	if err != nil {
		fmt.Println("数据库信息更新失败, err:", err)
		tx.Rollback()
		http.Redirect(w, r, "/", 302)
		return
	}

	// 获取受影响的行数
	num, err := db_res.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响行数失败, err:", err)

	} else {
		fmt.Println("此次修改受影响的行数: ", num)
	}

	http.Redirect(w, r, "/", 302)
}

// 取消点赞
func cancel_scraoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("取消点赞申请接入")
	r.ParseForm()

	article_id := r.Form.Get("ID")
	user_id := r.Form.Get("uid")

	fmt.Println(article_id, user_id)

	// 将该用户id从voted:表中去除
	// 减少source:表中对应文章的点赞分数
	// 减少article:表中积分的点赞分数
	// 修改数据库中的文章的点赞分数


	http.Redirect(w, r, "/", 302)
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Println("MySQL连接失败, err:", err)
		return
	}
	defer db.Close()
	fmt.Println("MySQL连接成功")

	if err := initRedis(); err != nil {
		fmt.Println("Redis连接失败, err:", err)
		return
	}

	fmt.Println("Redis连接成功")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/scrap", scrapHandler)
	http.HandleFunc("/cancel_scrao", cancel_scraoHandler)
	http.ListenAndServe(":80", nil)
}
