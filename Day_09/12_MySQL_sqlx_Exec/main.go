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

// 插入数据
func insertRowDemo() {
	sql := "insert into user(name, age) values(?,?)"
	// 执行sql
	ret, err := DB.Exec(sql, "T1", 18)
	if err != nil {
		fmt.Println("插入数据失败-1, err:", err)
		return
	}

	// 获取刚插入数据库信息的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("插入数据失败-1, err:", err)
		return
	}
	fmt.Println(id)
}

// 更新数据
func updateRowDemo() {
	sql := "update user set name=?,age=? where id=?"
	ret, err := DB.Exec(sql, "Smurfs_6", 23, 6)
	if err != nil {
		fmt.Println("更新数据失败, err:", err)
		return
	}

	// 获取受影响行数
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响行数失败, err:", err)
		return
	}
	fmt.Println(num)
}

// 删除数据
func deleteRowDemo() {
	sql := "delete from user where id=?"
	ret, err := DB.Exec(sql, 6)
	if err != nil {
		fmt.Println("删除数据失败, err:", err)
		return
	}

	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响行数失败, err:", err)
		return
	}
	fmt.Println(num)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err", err)
		return
	}

	fmt.Println("数据库连接成功")

	// insertRowDemo()

	// updateRowDemo()

	deleteRowDemo()
}
