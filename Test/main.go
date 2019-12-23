package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./index.html")
	if err != nil {
		fmt.Println("创建模板失败, err:", err)
		return
	}

	jsStr := "<script>alert('嘿嘿嘿')</script>"
	err = tmpl.Execute(w, jsStr)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
