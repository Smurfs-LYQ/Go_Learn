package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，检查MySQL数据库驱动是否正常，但不检查数据库信息
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试创建数据库连接
	err = DB.Ping()
	if err != nil {
		DB.Close()
		return err
	}

	return
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	// 数据库删除demo
	sql := "delete from user where id > ?"
	res, err := DB.Exec(sql, 5)
	if err != nil {
		fmt.Println("数据删除失败, err:", err)
		return
	}

	// 获取删除数据受影响行数
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取删除数据受影响行数失败, err:", err)
	}

	fmt.Printf("删除数据受影响行数: %d\n", n)
}
