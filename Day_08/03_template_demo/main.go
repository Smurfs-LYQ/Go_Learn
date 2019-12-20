package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 打开一个模板文件
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("文件加载失败, err:", err)
		return
	}

	// 用数据去渲染模板
	t.Execute(w, "我的世界")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
