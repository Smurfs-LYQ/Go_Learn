package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./index.html")
	if err != nil {
		panic(fmt.Sprintf("文件读取失败, err:%s", err))
	}
	w.Write(str)
}

func index(w http.ResponseWriter, r *http.Request) {
	// w: 代表了跟返回相关的所有内容
	// r: 代表了跟请求相关的所有内容

	// r.Method 可以返回请求的方法 (POST/GET)
	fmt.Println("请求的方法: ", r.Method)
	// 解析Form表单的数据
	r.ParseForm()

	// 获取Form表单的全部数据
	// fmt.Println(r.Form)
	w.Write([]byte(fmt.Sprintln(r.Form)))

	// 获取指定标签
	w.Write([]byte(r.Form.Get("username"))) // post和get发送过来的数据都可以接收到
	// w.Write([]byte(r.PostFormValue("username")))
	// w.Write([]byte(r.FormValue("password")))
	// w.Write([]byte(r.FormFile("photo")))
}

func main() {
	http.HandleFunc("/web", demo)
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
