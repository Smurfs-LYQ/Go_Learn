package main

// sqlx 示例 连接数据库

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

// DB 全局数据库连接对象(内置连接池)
var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	// 连接数据库，执行了open和ping的操作
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		DB.Close()
		return err
	}

	defer DB.Close()

	// 设置数据库最大连接数
	DB.SetMaxOpenConns(10)
	// 设置数据库最大空闲连接数
	DB.SetMaxIdleConns(5)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")
}
