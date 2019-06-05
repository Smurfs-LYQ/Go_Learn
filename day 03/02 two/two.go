// Golang整数类型的使用细节
package main

import(
	"fmt"
	"unsafe"
)

func main() {
	// 如何查看一个变量的数据类型以及占用的字节数
	var T1 int8 = 100
	fmt.Printf("T1的类型为%T T1占用的字节数为%d", T1, unsafe.Sizeof(T1))
	// fmt.Printf() // 可以用于格式化输出
	fmt.Println()

	// 检查一个变量占用的字节数
	fmt.Println("T1占用的数据类型为：", unsafe.Sizeof(T1))
	// unsafe.Sizeof(T1) // Sizeof是unsafe的一个方法，可以返回T1变量占用的字节数
}