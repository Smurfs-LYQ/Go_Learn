package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，检查数据库驱动
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试创建数据库连接，校验数据库信息
	err = DB.Ping()
	if err != nil {
		DB.Close()
		return err
	}

	return
}

// 事务demo
func transactionDemo() {
	tx, err := DB.Begin() // 事务开始
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Println("事务开始失败, err:", err)
		return
	}

	// 想要执行的操作-1
	sql_1 := "update user set name=?, age=age+? where id=?"
	_, err = tx.Exec(sql_1, "Smurfs_6", 2, 11)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("更新数据失败-1, err:", err)
		return
	}

	// 想要执行的操作-2
	sql_2 := "update user set name=?, age=age-? where id=?"
	// sql_2 := "update from user set name=?, age=age-? where id=?" // 有问题语句 会触发事务
	_, err = tx.Exec(sql_2, "Smurfs_7", 2, 12)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("更新数据失败-2, err:", err)
		return
	}

	err = tx.Commit() // 提交事务 代表所有修改语句都执行了
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println("事务提交失败, err:", err)
		return
	}
	fmt.Println("数据库修改完成")
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"

	err := initDB(dsn)
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	transactionDemo()
}
