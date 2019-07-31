package main

import "fmt"

/*
	1. 切片: 切片不能单独存在, 切片是数组的一个引用, 因此切片是引用类型
	2. 切片的长度可以改变, 因此, 切片是一个可变的数组
	3. 切片遍历方式和数组一样, 可以用len()求长度
	4. cap可以求出slice最大的容量, 0 <= len(slice) <= cap(array), 其中array是slice引用的数组
	5. 切片的定义, var 变量名 []类型, 比如: "var str []string" "var arr []int"
*/

func testSlice() {
	// 声明一个切片
	var T1 = []string{"hello", "world", "hello", "golang"}
	fmt.Println(T1)

	// 先声明一个数组, 再进行切片
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice []int = arr[2:5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 切片初始化: var slice []int = arr[start:end] 包含start到end之间的元素, 但不包含end
	// var slice []int = arr[0:end] 可以简写为 var slice []int=arr[:end]
	// var slice []int = arr[start:len(arr)] 可以简写为 var slice []int = arr[start:]
	// var slice []int = arr[0, len(arr)] 可以简写为 var slice []int = arr[:]
	// 如果要去掉切片最后一个元素, 可以这么写: slice = slice[:len(slice)-1]
}

func main() {
	// fmt.Println("123")
	testSlice()
}
