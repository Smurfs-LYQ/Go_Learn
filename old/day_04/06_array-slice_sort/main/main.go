package main

import (
	"fmt"
	"sort"
)

// 数组和切片的排序和查找

/*
	1. 排序操作主要都在sort包中, 导入就可以使用了
	2. 排序
		sort.Ints对整数进行排序
		sort.Strings对字符串进行排序
		sort.Float64s对浮点数进行排序
	3. 查找
		sort.SearchInts(a []int, b int) 从数组a中查找b
		sort.SearchFloats(a []float64, b float64) 从数组a中查找b
		sort.SearchStrings(a []string, b string) 从数组a中查找b
		注: 前提是数组a必须是有序, 因为在查找的时候会自动先排序再查找
*/

func testSortInts() {
	var a = make([]int, 5)   // 创建一个数组
	a = []int{1, 2, 5, 3, 4} // 给数组赋值
	fmt.Println("排序前: ", a)
	sort.Ints(a[:]) // 因为数组是值类型, 所以虽然已经排序好了, 但是a本身还是没有改变的, 改变的是a的调用复制体, 所以需要传一个切片进去, 因为切片是引用类型
	fmt.Println("排序后: ", a)

	// 查找指定值
	fmt.Println("#####查找#####")
	fmt.Println("4的位置在: ", sort.SearchInts(a, 4))
}

func testSortStrings() {
	var a = [5]string{"1", "2", "4", "5", "3"}
	fmt.Println("排序前: ", a)
	sort.Strings(a[:]) // 因为数组是值类型, 所以虽然已经排序好了, 但是a本身还是没有改变的, 改变的是a的调用复制体, 所以需要传一个切片进去, 因为切片是引用类型
	fmt.Println("排序后: ", a)

	// 查找指定值
	fmt.Println("#####查找#####")
	fmt.Println("2的位置在: ", sort.SearchStrings(a[:], "2"))

}

func testSortFloat64s() {
	a := [5]float64{0.1, 0.2, 0.5, 0.3, 0.4}
	fmt.Println("排序前: ", a)
	sort.Float64s(a[:])
	fmt.Println("排序后: ", a)

	// 查找指定值
	fmt.Println("#####查找#####")
	fmt.Println("0.5的位置在: ", sort.SearchFloat64s(a[:], 0.5))

}

func main() {
	fmt.Println("Sort.Ints")
	testSortInts()
	fmt.Println("Sort.Strings")
	testSortStrings()
	fmt.Println("Sort.Float64s")
	testSortFloat64s()
}
