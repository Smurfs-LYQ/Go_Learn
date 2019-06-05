// golang整数类型的使用
package main

import "fmt"

func main() {
	var one int = 10
	fmt.Println(one)

	// 测试一下int8的范围 -128~127
	var i int8 = -128
	fmt.Println(i)

	// 测试一下uint8的范围 0~255
	var ui uint8 = 255
	fmt.Println(ui)

	// 其他的int16、int32、int64、uint16、uint32、uint64类推

	// int、uint、rune、byte的使用
	var T1 int = 8900
	fmt.Println(T1)
	var T2 uint = 0
	fmt.Println(T2)
	var T3 byte = 255
	fmt.Println(T3)
}