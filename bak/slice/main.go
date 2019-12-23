// slice 切片
package main

import (
	"fmt"
)

// 介绍：
	// slice是一种变长的数组，其数据结构中有指向数组的指针，所以是一种引用类型

func main() {
	// type slice struct {
	// 	arr_1b unsafe.Pointer
	// 	len int
	// 	cap int
	// }

	// 由数组创建的切片
	var arr_1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr_1[0:3]
	s2 := arr_1[:3]
	s3 := arr_1[3:]
	fmt.Printf("%v\t%T\n", s1, s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("#################")

	// 通过内置函数make创建切片
	// len=10, cap=10
	a := make([]int, 10)		// 0 0 0 0 0 0 0 0 0 0
	// len=10, cap=15
	b := make([]int, 10, 15)	// 0 0 0 0 0 0 0 0 0 0
	
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\n", b)

	fmt.Println("#################")

	// 实例
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s4 := make([]int, 2, 4)
	c1 := a1[0:3]

	// 检测切片长度
	fmt.Println(len(s4))
	// 检测切片底层数组容量
	fmt.Println(cap(s4))
	// 对切片追加元素
	s4 = append(s4, 1)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	
	fmt.Println("#################")

	s4 = append(s4, c1...) //后面的...是规定的格式
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	fmt.Println("#################")

	d := make([]int, 2, 2)
	// 把c中的内容拷贝到d中，只拷贝最小的，并且会遵守d设置的长度限制
	copy(d, c1)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
}