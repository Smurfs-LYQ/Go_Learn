package main

import "fmt"

/*
	1. 数组: 是同一种数据类型的固定长度序列
	2. 数组定义: var a[len]int 比如: var a[5]int
	3. 长度是数组类型的一部分, 因此, "var a[5]int"和"var a[10]int"是不同的类型
	4. 数组可以通过下标进行访问, 下标是从0开始的, 最后一个元素下标是: len-1
	5. 访问越界: 如果下标在数组合法范围之外, 则触发访问越界, 会painc
	6. 因为数组是值类型, 所以复制的数组被修改不会影响原数组
*/

func main() {
	var a [10]int // 定义一个10个元素的数组
	// a[0] = 0 //通过下标给指定元素赋值

	// 通过for循环给指定元素赋值或访问
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	// 通过for range获取下标访问或赋值指定元素
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 数组初始化
	// var age [5]int = [5]int{1, 2, 3} //如果数组的元素没有满足设定的数组长度, 则使用对应类型的默认值
	// var age = [5]int{1, 2, 3, 4, 5} // 相比上一种声明方式在前半部分省去了数据类型
	// var age = [...]int{1, 2, 3, 4, 5} // 任意长度的数组
	// var age = [5]string{3: "hello world", 4: "tom"} // 初始化赋值的时候通过下标给指定元素设置值
	// fmt.Println(age)

	//多维数组
	var age0 [2][2]string = [2][2]string{{"hello", "world"}, {"hello", "golang"}} // 二维数组
	var age1 [3][2]int = [...][2]int{{1, 2}, {3, 4}, {5, 6}}                      // 三维数组
	fmt.Println(age0)
	fmt.Println(age1)

	for _, v_0 := range age0 {
		fmt.Println(v_0)
		for _, v_1 := range v_0 {
			fmt.Println(v_1)
		}
	}
}
