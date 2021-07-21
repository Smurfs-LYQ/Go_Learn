package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

var db *sqlx.DB

func initDB() (err error) {
	// dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test?parsetime=true" // parsetime : 如果连接的时候有与时间相关的操作，加上这个Go会自动帮忙转换一下
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(15)

	return
}

// 注册用户
func createUser(username, password string) (status int) {
	// 判断此用户是否存在
	sql := "select count(id) from user where name=?"
	var num int
	err := db.Get(&num, sql, username)
	if err != nil {
		fmt.Println(err)
	}

	if num > 0 {
		return 3 // 用户已存在
	}

	// 用户不存在插入数据库
	sql = "insert into user(name, password) values(?,?)"
	_, err = db.Exec(sql, username, password)
	if err != nil {
		fmt.Println(err)
		return 2 // 用户入库失败
	}

	return 1 // 用户入库成功
}

func loginUser(username, password string) (status int) {
	sql := "select password from user where name=?"
	var pwd string
	err := db.Get(&pwd, sql, username)
	if err != nil || pwd != password {
		return 2 // 登录失败
	}

	return 1 // 登录成功
}
