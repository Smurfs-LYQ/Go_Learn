package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 用户注册
func register(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/register.html")
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		w.WriteHeader(500) // 效果同上
	}

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		// fmt.Println(username, password)
		status := createUser(username, password)
		switch status {
		case 1:
			// 用户入库成功
			http.Redirect(w, r, "/login", 302)
		case 2:
			// 用户入库失败
			t.Execute(w, "注册失败")
		case 3:
			// 该用户已存在
			t.Execute(w, "该用户名已被占用")
		}
	} else {
		t.Execute(w, nil)
	}
}

// 用户登录
func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/login.html")
	if err != nil {
		w.WriteHeader(500)
	}

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		status := loginUser(username, password)
		switch status {
		case 1:
			http.Redirect(w, r, "https://www.baidu.com", 302)
		case 2:
			t.Execute(w, "登录失败")
		}
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic("数据库连接失败")
	}

	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("启动http server失败: ", err)
		return
	}
}
