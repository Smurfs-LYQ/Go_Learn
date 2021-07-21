package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 增
func addHandler(c *gin.Context) {
	msg := c.Query("msg")
	if msg == "" {
		c.HTML(http.StatusOK, "templates/add.html", nil)
	} else {
		c.HTML(http.StatusOK, "templates/add.html", gin.H{
			"msg": msg,
		})
	}
}

func addDoHandler(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	price, err := strconv.ParseFloat(strings.TrimSpace(c.PostForm("price")), 64)
	if err != nil {
		fmt.Printf("价格转Int类型失败, err:%s", err)
		c.Redirect(http.StatusMovedPermanently, "/book_add?msg=price")
	}

	err = addBookDB(name, price)
	if err != nil {
		fmt.Printf("数据库插入失败, err:%s", err)
		c.Redirect(http.StatusMovedPermanently, "/book_add?msg=mysql")
	}

	c.Redirect(http.StatusMovedPermanently, "/book_lists")
}

// 删
func delHandler(c *gin.Context) {
	id, err := strconv.Atoi(strings.TrimSpace(c.Query("id")))
	if err != nil {
		fmt.Println("ID传递异常")
		c.Redirect(http.StatusMovedPermanently, "/book_lists?msg=删除失败")
	}

	err = delBookDB(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/book_lists?msg=删除失败")
	}

	c.Redirect(http.StatusMovedPermanently, "/book_lists")
}

// 改
func editHandler(c *gin.Context) {
	id, err := strconv.Atoi(strings.TrimSpace(c.Query("id")))
	if err != nil {
		fmt.Println("ID传递异常")
		c.Redirect(http.StatusMovedPermanently, "/book_lists?msg=无法编辑")
	}

	data, err := editBookDB_Get(id)
	if err != nil {
		msg := fmt.Errorf("sql: no rows in result set")
		if err == msg {
			c.Redirect(http.StatusMovedPermanently, "/book_lists?msg=没有这本书")
		}
		c.Redirect(http.StatusMovedPermanently, "/book_lists?msg=修改失败")
	}

	c.HTML(http.StatusOK, "templates/edit.html", gin.H{
		"id":    id,
		"title": data.Title,
		"price": data.Price,
		"msg":   c.Query("msg"),
	})
}

func editDoHandler(c *gin.Context) {
	id := strings.TrimSpace(c.PostForm("id"))

	var (
		name  string
		price float64
		err   error
	)
	if strings.TrimSpace(c.PostForm("name")) != "" {
		name = strings.TrimSpace(c.PostForm("name"))
	} else {
		name = strings.TrimSpace(c.PostForm("bak_name"))
	}
	if strings.TrimSpace(c.PostForm("price")) != "" {
		price, err = strconv.ParseFloat(strings.TrimSpace(c.PostForm("price")), 64)
	} else {
		price, err = strconv.ParseFloat(strings.TrimSpace(c.PostForm("bak_price")), 64)
	}

	if err != nil {
		fmt.Printf("价格转Int类型失败, err:%s", err)
		c.Redirect(http.StatusMovedPermanently, "/book_add?msg=price")
	}

	err = editBookDB(id, name, price)
	if err != nil {
		fmt.Printf("数据库修改失败, err:%s", err)
		c.Redirect(http.StatusMovedPermanently, "/book_edit?msg=修改失败")
	}

	c.Redirect(http.StatusMovedPermanently, "/book_lists")
}

// 查
func listHandler(c *gin.Context) {
	msg := c.Query("msg")

	data, err := listBookDB()
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "templates/list.html", gin.H{
		"data": data,
		"msg":  msg,
	})
}

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 增加书籍
	r.GET("/book_add", addHandler)
	r.POST("/book_add", addDoHandler)
	// 删除书籍
	r.GET("/book_del", delHandler)
	// 修改书籍
	r.GET("/book_edit", editHandler)
	r.POST("/book_edit", editDoHandler)
	// 书籍列表
	r.GET("/book_lists", listHandler)

	r.Run()
}
