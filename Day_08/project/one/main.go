package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	使用net/http编写一个server端
	返回一个html登录页面，登录页面要有用户名和密码
	点提交能够把表单的数据发送到server端，server端要校验用户名(smurfs)和密码(test)是否正确
*/

func index(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadFile("./one.html")
	if err != nil {
		panic(fmt.Sprintf("页面加载失败, err:", err))
	}
	w.Write(res)
}

func login(w http.ResponseWriter, r *http.Request) {
	// 解析Form表单的数据
	r.ParseForm()

	// 打印Form表单中的所有数据
	w.Write([]byte(fmt.Sprintln(r.Form)))

	// 获取用户名和密码
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// 判断
	if username == "smurfs" && password == "test" {
		w.Write([]byte("登录成功"))
	} else {
		w.Write([]byte("登录失败"))
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
