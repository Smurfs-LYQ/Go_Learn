package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Moto struct {
	Name         string
	Engine       string
	Displacement int
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("页面加载失败, err:", err)
		return
	}

	MotoMap := map[int]Moto{
		1: {"Ducati V4R", "v4", 999},
		2: {"Aprilia RSV4", "v4", 999},
		3: {"KTM RC9", "V2", 1199},
	}

	t.Execute(w, MotoMap)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
