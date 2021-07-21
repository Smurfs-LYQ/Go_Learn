package main

import (
	"database/sql"
	"fmt"

	// 注册驱动
	_ "github.com/go-sql-driver/mysql"
)

// 使用连接池方式连接MySQL

// DB 数据库连接句柄, 全局变量
var DB *sql.DB

// initDB 初始化数据库连接
func initDB(dsn string) (err error) {
	// 打开数据库，检查有没有数据库驱动，但不会校验账号密码是否正确
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// defer DB.Close()
	// 尝试与数据库建立连接(校验dsn)
	err = DB.Ping()
	if err != nil {
		DB.Close()
		return err
	}

	// 设置最大连接数 100个
	DB.SetMaxOpenConns(10)
	// 设置最大空闲连接数 20个
	DB.SetMaxIdleConns(5)

	return
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")

	for i := 0; i < 100; i++ {
		var ch1 chan string
		ch1 = make(chan string, 1)
		go test(ch1)
		fmt.Println(i, <-ch1)
	}
}

func test(ch1 chan string) {
	err := DB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	ch1 <- "数据库连接成功"
}
