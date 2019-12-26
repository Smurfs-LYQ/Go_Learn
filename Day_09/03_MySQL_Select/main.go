package main

import (
	// 注册驱动
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB(dsn string) (err error) {
	// 打开数据库，检查有没有数据库驱动，但不会校验账号密码是否正确
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立链接
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单行数据
func queryRowDemo(sqlStr string) {
	/*
		row := DB.QueryRow(sqlStr, 1)             // sql语句，对应sql语句的可变参数(如果没有可不填)
		err = row.Scan(&user.id, &user.name, &user.age)		  // 需要将接收查询结果的变量传入
	*/
	err := DB.QueryRow(sqlStr, 1).Scan(&user.id, &user.name, &user.age) // 效果同上
	if err != nil {
		fmt.Println("查询失败, err:", err)
		return
	}

	var user = user{}

	fmt.Println(user)
}

// 查询多条数据
func queryMultiDemo(sqlStr string) {
	rows, err := DB.Query(sqlStr, 1)
	if err != nil {
		fmt.Println("查询失败, err", err)
		return
	}
	// 函数结束释放rows, 如果不释放就会一直占取连接数
	defer rows.Close()

	var user = user{}

	// 循环读取数据
	for rows.Next() { // 一直拿取rows中的数据，只要rows中有数据，rows.Next()就会一直返回true
		rows.Scan(&user.id, &user.name, &user.age)
		fmt.Println(user)
	}
}

func main() {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test" // 用户名:密码@连接方式(IP:端口)/库名
	err := initDB(dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err)
		return
	}

	fmt.Println("数据库连接成功")

	// CRUD
	sqlStr := "select * from user where id>=?" // ?是占位符

	// 查询单条数据
	fmt.Println("查询单条数据")
	queryRowDemo(sqlStr)

	// 查询多条数据
	fmt.Println("查询多条数据")
	queryMultiDemo(sqlStr)
}
