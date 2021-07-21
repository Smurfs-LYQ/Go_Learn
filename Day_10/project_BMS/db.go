package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := ("root:smurfs@tcp(127.0.0.1:3306)/go_test")

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return
}

// 添加数据
func addBookDB(name string, price float64) (err error) {
	sql := "insert into book(title, price) values(?,?)"

	_, err = db.Exec(sql, name, price)
	if err != nil {
		return err
	}
	return
}

// 查询数据
func listBookDB() (data []book, err error) {
	sql := "select id,title,price from book"

	err = db.Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 查询要修改数据页面
func editBookDB_Get(id int) (data book, err error) {
	sql := "select title,price from book where id=?"
	err = db.Get(&data, sql, id)
	fmt.Println(data, err)
	if err != nil {
		return
	}

	return
}

func editBookDB(id string, name string, price float64) (err error) {
	sql := "update book set title=?, price=? where id=?"
	_, err = db.Exec(sql, name, price, id)
	if err != nil {
		return err
	}

	return
}

func delBookDB(id int) (err error) {
	sql := "delete from book where id=?"
	_, err = db.Exec(sql, id)
	if err != nil {
		return err
	}
	return
}
