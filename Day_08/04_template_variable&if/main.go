package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	UserName string
	Age      int
	Sex      int
}

func index(w http.ResponseWriter, r *http.Request) {
	// 打开一个模板文件
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	user := User{
		UserName: "Smurfs",
		Age:      18,
		Sex:      1,
	}

	// 用数据渲染模板
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
