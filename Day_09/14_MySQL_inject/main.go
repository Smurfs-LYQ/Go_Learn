package main

// SQL注入问题示例

import (
	// "fmt"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	return
}

// sql注入问题示例
func sqlInjectDemo(name string) {
	sql := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL: %s\n", sql)

	// DB.Select(sql)
	var u []user
	err := DB.Select(&u, sql)
	if err != nil {
		fmt.Println("查询失败, err:", err)
		return
	}
	for _, u := range u {
		fmt.Println(u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	// 正常查询
	// select id, name, age from user where name='Smurfs_1'
	// sqlInjectDemo("Smurfs_1")

	// 不正常查询_1
	/*
		select id, name, age from user where name='Smurfs_1' or 1=1#'
		解析:
			Smurfs_1' or 1=1#
			无论 "Smurfs_1" 被任何值替代都没关系，因为他查出来的数据无关紧要, 并且 ' 刚好也截止了name的赋值操作
			重点在于 or 1=1，这条语句代表着where后面的条件跟 1=1 这个条件成立一个就可以了，并且1=1永远等于true
			# 是注释，屏蔽了后面剩下的所有SQL语句，加上 1=1 永远成立，相当于这条SQL语句没有where条件 等同于 "select id, name, age from user"
	*/
	// sqlInjectDemo("Smurfs_1' or 1=1#")

	// 不正常查询_2
	/*
		select id, name, age from user where name='xxx' union select * from user #'
		解析:
			xxx' union select * from user #
			无论 "xxx" 被任何值替代都没关系，因为他查出来的数据无关紧要, 并且 ' 刚好也截止了name的赋值操作
			union 的作用是联合查询，所以后面跟的语句可以任由其调整
			# 是注释，屏蔽了后面剩下的所有SQL语句
	*/
	// sqlInjectDemo("xxx' union select * from user #")

	// 不正常查询_3
	sqlInjectDemo("xxx' and select count(*) from user #")
}
