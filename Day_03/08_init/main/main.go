package main

import "fmt"

func init() {
	// 程序开始时自动执行
	fmt.Println("这个是init函数")
}

func main() {
	fmt.Println("这个是main函数")
}
