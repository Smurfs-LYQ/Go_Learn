package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Moto struct {
	MotoName string
	Age      int
	Sex      int
}

func index(w http.ResponseWriter, r *http.Request) {
	// 打开一个模板文件
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("文件打开失败, err:", err)
	}

	/*
		motoSlice := []Moto{
			{"Ducati", 18, 1},
			{"KTM", 19, 1},
			{"Aprilia", 20, 0},
		}
		t.Execute(w, motoSlice)
	*/

	motoMap := map[int]Moto{
		1: {"Kawasaki", 18, 1},
		2: {"Yamaha", 19, 1},
		3: {"Honda", 20, 0},
		4: {"Suzuki", 21, 1},
	}

	t.Execute(w, motoMap)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
