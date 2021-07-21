package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html", "./europe.html")
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
