package main

import "fmt"

type T1 struct {
	name   string
	int    // 匿名字段
	string // 匿名字段
}

func main() {
	var Test_1 = T1{
		name:   "Smurfs",
		int:    21,  // 匿名字段赋值
		string: "男", // 匿名字段赋值
	}

	fmt.Println(Test_1)
}
