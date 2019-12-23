package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	使用net/http编写一个server端
	返回一个html登录页面，登录页面要有用户名和密码
	点提交能够把表单的数据发送到server端，server端要校验用户名(smurfs)和密码(test)是否正确
*/

func login(w http.ResponseWriter, r *http.Request) {
	// 创建模板
	t, err := template.ParseFiles("./one.html")
	if err != nil {
		panic(fmt.Sprintf("页面加载失败, err:", err))
	}

	// 判断用户是获取页面还是请求登录
	if r.Method == "POST" {
		// 解析Form表单的数据
		r.ParseForm()

		// 获取用户名和密码
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 判断
		if username == "smurfs" && password == "test" {
			// 登录成功，跳转页面
			http.Redirect(w, r, "https://www.baidu.com", 302)
		} else {
			t.Execute(w, "用户名或密码错误")
		}
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
