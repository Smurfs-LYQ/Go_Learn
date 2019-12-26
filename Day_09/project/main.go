package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
	age  int
}

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，测试数据库驱动
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试建立数据库连接, 校验数据库信息
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

	sql := "select * from user where name=?"
	rows, err := DB.Query(sql, "Smurfs_5")
	if err != nil {
		fmt.Println(err)
		return
	}
	var u user
	for rows.Next() {
		rows.Scan(&u.id, &u.name, &u.age)
		fmt.Println(u)
	}
}
