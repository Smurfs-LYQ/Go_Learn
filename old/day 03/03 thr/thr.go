// Golang浮点型使用
package main

import(
	"fmt"
	"unsafe"
)

func main() {
	// 定义一个浮点型变量
	var T1 float32 = -0.00035
	var T2 float64 = -23562341.23412
	fmt.Println("T1=", T1, "T2=", T2)
	fmt.Printf("T1数据类型为 %T\n", T1)
	fmt.Println("T1占用了", unsafe.Sizeof(T1), "位字节")

	// 尾数部分可能丢失，造成精度损失。-123.0000901
	var T3 float32 = -123.0000901
	var T4 float64 = -123.0000901
	fmt.Println(T3)
	fmt.Println(T4)
}