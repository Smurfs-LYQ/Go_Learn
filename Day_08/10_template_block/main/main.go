package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 按照正则匹配规则解析模板文件
	t, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	// 渲染指定的模板文件
	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("页面渲染失败, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
