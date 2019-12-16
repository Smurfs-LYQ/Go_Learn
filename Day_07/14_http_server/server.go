package main

import (
	"fmt"
	"net/http" // 专门为http协议写的包
)

// HTTP Server

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='color:red;'>hello world!</h1>"))
}

func main() {
	http.HandleFunc("/", sayHello)           // 注册路由: 当你访问 / 这个路径的时候就执行sayHello函数
	err := http.ListenAndServe(":9090", nil) // 建立监听
	if err != nil {
		fmt.Println("服务启动失败, err: ", err)
		return
	}
}
