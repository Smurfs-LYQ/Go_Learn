package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	return
}

func transDemo() {
	// 开启事务
	tx, err := DB.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("开启事务失败, err:", err)
		return
	}

	// 数据库操作
	sql_1 := "update user set name=?,age=age-? where id=?"
	tx.MustExec(sql_1, "smurfs_6", 2, 7) // 第三方库带Must的一般出现问题都会panic
	// sql_2 := "update user set name=?,age=age+? where id=?" // 错误语句，会触发事务
	sql_2 := "update from user set name=?,age=age+? where id=?" // 错误语句，会触发事务
	tx.MustExec(sql_2, "smurfs_7", 2, 8)

	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback()
		fmt.Println("修改数据库失败, err:", err)
		return
	}
	fmt.Println("更新数据库成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	transDemo()
}
