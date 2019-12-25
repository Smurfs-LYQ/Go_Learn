package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，检查是否有mysql数据库驱动，但并不检查数据库信息
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

	// 数据库更新数据demo
	sql := "update user set name=? where id=?"
	res, err := DB.Exec(sql, "Smurfs_5", 5)
	if err != nil {
		fmt.Println("数据库信息更新失败, err:", err)
		return
	}

	// 获取受影响行数
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取更新信息受影响行数失败, err", err)
	}
	fmt.Printf("更新信息受影响行数为: %d\n", n)
}
