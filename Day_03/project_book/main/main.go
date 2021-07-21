package main

import (
	"Go_Learn/Day_03/project_book/books"
	"fmt"
	"os"
)

func Manager() {
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
	fmt.Println()
}

func main() {
	/*
		使用函数实现一个图书管理系统
		每本书有 书名、作者、价格、上架信息
		用户可以在控制台 添加书籍、修改书籍信息、打印所有的书籍列表
	*/
	for {
		Manager()

		var do int
		fmt.Printf("%s", "请输入操作指令: ")
		fmt.Scanln(&do)
		switch do {
		case 1:
			books.Add()
		case 2:
			books.Set()
		case 3:
			books.Show()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("打印书籍列表")
		}
	}
}
