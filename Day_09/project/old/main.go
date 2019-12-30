package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"github.com/jmoiron/sqlx"          // 导入第三方标志库
)

type user struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:smurfs@tcp(127.0.0.1:3306)/go_test"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		db.Close()
	}

	return
}

func mysql_login(name, pwd string) (state int) {
	sql := "select id, name, password from user where name=?"

	var u user
	err := db.Get(&u, sql, name)
	if err != nil {
		return 3 // 3 代表没有这个用户
	}

	if u.Password != pwd {
		return 2 // 2 代表密码错误
	}

	return 1 // 1 代表登录成功
}

func mysql_sign(name, pwd string) (state int) {
	sql := "select name from user where name=?"
	var u user
	err := db.Get(&u, sql, name)
	if err == nil {
		return 3 // 用户名已存在
	}

	sql = "insert into user(name, password) values(?, ?)"

	_, err = db.Exec(sql, name, pwd)
	if err != nil {
		fmt.Println("插入数据失败, err:", err)
		return 2 // 创建用户失败
	}

	return 1 // 创建用户成功
}

func login(w http.ResponseWriter, r *http.Request) {
	// 导入文件模板
	tpl, err := template.ParseFiles("./html/login.html")
	if err != nil {
		fmt.Println("文件加载失败, err:", err)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		state := mysql_login(username, password)
		switch state {
		case 1:
			http.Redirect(w, r, "https://www.baidu.com", 302)
		case 2:
			tpl.Execute(w, "密码错误")
		case 3:
			tpl.Execute(w, "该用户不存在")
		}

	} else {
		// 渲染模板
		tpl.Execute(w, nil)
	}
}

func sign_in(w http.ResponseWriter, r *http.Request) {
	// 导入文件模板
	tpl, err := template.ParseFiles("./html/sign_in.html")
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")
		state := mysql_sign(username, password)
		fmt.Println(username, password, state)
		switch state {
		case 1:
			tpl.Execute(w, "用户注册成功")
		case 2:
			tpl.Execute(w, "用户注册失败")
		case 3:
			tpl.Execute(w, "该用户已存在")
		default:
			tpl.Execute(w, nil)
		}
	} else {
		tpl.Execute(w, nil)
	}
}

func main() {
	// 启动数据库连接
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败, err:", err)
		return
	}

	fmt.Println("数据库连接成功")

	http.HandleFunc("/login", login)
	http.HandleFunc("/sign_in", sign_in)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
