package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取环境变量
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("GOPATH"))
}
