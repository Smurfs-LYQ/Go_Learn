package main

import "fmt"

func add(a, b float64) float64 { // func 函数名(a 数据类型, b 数据类型) 返回值数据类型 {
	return float64(a) + float64(b) // 函数体
} // 函数尾

func main() {
	fmt.Println(add(1, 1.1))
}
