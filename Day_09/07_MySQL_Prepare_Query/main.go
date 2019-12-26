package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL预处理

type user struct {
	id   int
	name string
	age  int
}

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，校验数据库驱动
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立链接，校验数据库信息
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

// 预处理数据库查询-Query
func prepareQueryRowDemo() {
	sql := "select * from user where id=?"
	// 1. 将需要执行的SQL语句发送到MySQL服务端
	stmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println("prepare failed, err:", err)
		return
	}

	// 2. 把程序最后关闭的工作做好
	defer stmt.Close()

	// 3. 执行重复查询语句-查询id为10以内的所有user
	for i := 0; i < 10; i++ {
		var u user
		err := stmt.QueryRow(i).Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("查询失败, err:", err)
			continue
		}
		fmt.Println(u)
	}
}

// 预处理数据库查询-QueryRow
func prepareQueryDemo() {
	sql := "select * from user where name=?"
	// 1. 将需要执行的SQL语句发送到MySQL服务端
	stmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println("prefare faield, err:", err)
		return
	}

	// 2. 把程序最后关闭的工作做好
	defer stmt.Close()

	// 3. 执行重复查询语句-
	for i := 0; i <= 5; i++ {
		name := fmt.Sprintf("Smurfs_%d", i)
		fmt.Println(name)
		res, err := stmt.Query(name)
		if err != nil {
			fmt.Println("查询失败-1, err:", err)
			continue
		}

		defer res.Close()

		var u user
		for res.Next() {
			err := res.Scan(&u.id, &u.name, &u.age)
			if err != nil {
				fmt.Println("查询失败-2, err:", err)
				continue
			}
			fmt.Println(u)
		}
	}
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}
	fmt.Println("数据库连接成功")

	// prepareQueryRowDemo()

	prepareQueryDemo()
}
