package main

import "net/http"

import "fmt"

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("请求失败, err:", err)
	}

	fmt.Println(resp)
}
