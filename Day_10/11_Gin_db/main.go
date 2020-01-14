package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		DB.Close()
		return err
	}

	// defer DB.Close()

	// 设置最大连接数
	DB.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(5)

	return
}

type User struct {
	Id   int    `db:"user_id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Book struct {
	Id    int    `db:"book_id"`
	Title string `db:"title"`
	Price int    `db:"price"`
	User
}

func bookHandler(c *gin.Context) {
	var book []Book
	// 跨表查询
	sql := `select 
				book.id as book_id, book.title, book.price, user.id as user_id, user.name, user.age
			from book,user where
				book.user_id = user.id`
	err := DB.Select(&book, sql)
	if err != nil {
		fmt.Println("查询失败", err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询失败",
			"err": err,
		})
	}

	// for _, v := range book {
	// 	fmt.Println(v)
	// }
	c.JSON(http.StatusOK, gin.H{
		"msg": book,
	})
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("book_list", bookHandler)

	r.Run()
}
