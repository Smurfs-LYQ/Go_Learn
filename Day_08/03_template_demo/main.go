package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("文件加载失败, err:", err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入w
	t.Execute(w, "我的世界")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
