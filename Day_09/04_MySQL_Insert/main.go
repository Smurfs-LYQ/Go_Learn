package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB 声明一个数据库操作的全局对象
var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，检查有没有数据库驱动，这一步并不会检查数据库用户名密码等信息
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	// 数据库信息插入
	sql := "insert into user(name, age) values(?, ?)"
	ret, err := DB.Exec(sql, "Smurfs", 22)
	if err != nil {
		fmt.Println("数据库信息插入失败, err:", err)
	}

	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("数据库信息插入失败, err:", err)
	}
	fmt.Printf("刚才插入数据库信息的ID为: %d\n", id)
}
