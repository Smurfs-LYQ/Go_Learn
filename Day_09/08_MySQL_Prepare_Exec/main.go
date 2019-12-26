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

	// 尝试与数据库建立连接，校验数据库信息
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func prepareDemo() {
	// insert
	// sql := "insert into user(name, age) values(?,?)"
	// update
	// sql := "update user set name=?, age='18' where name='smurfs'"
	// delete
	sql := "delete from user where id = ?"
	// 1. 把要执行的命令发送给MySQL服务端
	stmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println("prepare failed, err:", err)
		return
	}
	// 2. 把程序最后关闭准备工作做好
	defer stmt.Close()

	// 预处理的用法跟内置的方法跟正常操作的DB非常的相似, 查询使用Query/Queryrow, 增删改使用Exec等
	// 3. 执行重复的插入命令
	for i := 5; i < 10; i++ {
		// insert
		// res, err := stmt.Exec(fmt.Sprintf("Smurfs_%d", i), rand.Intn(25))
		// update
		// res, err := stmt.Exec(fmt.Sprintf("Smurfs_%d", i))
		// delete
		res, err := stmt.Exec(i)

		if err != nil {
			fmt.Println("数据库操作失败, err:", err)
			continue
		}

		/*
			// 获取最后插入数据库数据的ID
			id, err := res.LastInsertId()
			if err != nil {
				fmt.Println("获取ID失败, err:", err)
				continue
			}
			fmt.Println("最后一次插入数据的ID: ", id)
		*/

		// 获取受影响行数
		n, err := res.RowsAffected()
		if err != nil {
			fmt.Println("获取受影响行数失败, err:", err)
			continue
		}
		fmt.Println("受影响行数: ", n)
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

	prepareDemo()
}
