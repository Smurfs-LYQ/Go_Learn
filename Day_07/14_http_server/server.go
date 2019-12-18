package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // 专门为http协议写的包
)

// HTTP Server

func sayHello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1 style='color:red;'>hello world!</h1>"))
	str, err := ioutil.ReadFile("./test.html")
	if err != nil {
		fmt.Println("文件读取失败，err:", err)
		return
	}
	w.Write(str)
}

/*
	使用浏览器访问本地的9090端口
*/
func main() {
	http.HandleFunc("/", sayHello) // 注册路由: 当你访问 / 这个路径的时候就执行sayHello函数
	// err := http.ListenAndServe(":9090", nil) // 建立监听
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 建立监听
	if err != nil {
		fmt.Println("服务启动失败, err: ", err)
		return
	}
}
