package main

import (
	"database/sql"
	"fmt"
	// 只用到了这个包中的init(), init()注册了mysql数据库的驱动
	_ "github.com/go-sql-driver/mysql"
)

// database/sql 连接MySQL示例代码

func main() {
	// 数据源名字 dataSourceName
	// dsn:"user:password@tcp(ip:port)/databasename"
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"

	/*
		调用标准库的方法，打开数据库，前提是要注册对应数据库的驱动
		参数: 驱动名称，数据源信息
		注意: 这里并没有真正的去连数据库，只是检查有没有数据库驱动，但不会校验账号密码是否正确
	*/
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("打开数据库错误, err:", err)
		return
	}
	// 关闭数据库
	defer db.Close()

	// 尝试连接一下数据库，校验(dsn)用户名密码是否正确
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败, err:", err)
		return
	}
	fmt.Println("连接数据库成功")
}
