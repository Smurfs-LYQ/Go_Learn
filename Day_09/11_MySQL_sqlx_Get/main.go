package main

// sqlx 示例 查询

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

var DB *sqlx.DB

type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		DB.Close()
		return err
	}

	return
}

// 查询单行
func queryRowDemo() {
	sql := "select id,name,age from user where id=?"
	var u user
	err := DB.Get(&u, sql, 1)
	if err != nil {
		fmt.Println("查询失败, err:", err)
		return
	}
	fmt.Println(u)
}

// 查询多行
func queryMultiDemo() {
	sql := "select id,name,age from user where id > ?"
	var u []user
	err := DB.Select(&u, sql, 0)
	if err != nil {
		fmt.Println("查询失败, err:", err)
		return
	}
	for _, v := range u {
		fmt.Println(v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	// queryRowDemo()

	queryMultiDemo()
}
