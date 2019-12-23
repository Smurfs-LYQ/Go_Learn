// 数组
package main

import "fmt"

func main() {
	// 确定长度的数组
	arr_1 := [2]string{"LYQ", "男"}
	fmt.Println(arr_1)

	// 不确定长度的数组，由后面的初始化列表数量来确定其长度
	arr_2 := [...]int{1,2,3}
	fmt.Println(arr_2)

	// 确定长度的数组，通过索引值进行初始化，没有初始化元素时使用类型默认值，被跳过的索引使用类型的默认值
	arr_3 := [4]string{ 0:"Smurfs", 1:"男", 3:"备用参数" }
	fmt.Println(arr_3)

	// 不确定长度的数组，通过索引值进行初始化，数组长度由最后一个索引值确定，没有指定索引的元素被初始化为类型的零值，被跳过的索引使用类型的默认值
	arr_4 := [...]int{ 1:1, 2:2, 3:3, 5:4 }
	fmt.Println(arr_4)

	fmt.Println("###############")

	// 访问数组
	fmt.Println(arr_1[0])
	for i, v := range arr_1 { // range在这里的用处是返回数组的索引
		fmt.Print(i,":")
		fmt.Println(v)
	}

	fmt.Println("###############")
	
	// 查看数组的长度
	fmt.Println(len(arr_3))
	for i := 0; i < len(arr_3); i++ {
		fmt.Println(arr_3[i])
	}
}