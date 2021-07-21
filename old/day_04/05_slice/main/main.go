package main

import "fmt"

/*
	1. 切片: 切片不能单独存在, 切片是数组的一个引用, 因此切片是引用类型
	2. 切片的长度可以改变, 因此, 切片是一个可变的数组
	3. 切片遍历方式和数组一样, 可以用len()求长度
	4. cap可以求出slice最大的容量, 0 <= len(slice) <= cap(array), 其中array是slice引用的数组
	5. 切片的定义, var 变量名 []类型, 比如: "var str []string" "var arr []int"
	6. 切片容量是可变的, 它会根据容量的大小重新分配内存
	7. string类型的切片, 值是不可变的
*/

func testSlice() {
	// 声明一个切片
	var T1 = []string{"hello", "world", "hello", "golang"}
	fmt.Println(T1)

	// 通过make来创建切片
	/*
		var slice []type = make([]type, len)
		slice := make{[]type, cap}
		slice := make([]type, len, cap)
	*/
	var T2 []int = make([]int, 3) // 使用var声明一个切片
	T2 = []int{1, 2, 3}

	T3 := make([]int, 3) // 简短格式声明一个切片
	T3 = []int{4, 5, 6}

	T4 := make([]int, 3, 3) // 简短格式声明一个切片, 并声明其长度和容量
	T4 = []int{7, 8, 9}

	fmt.Println(T2)
	fmt.Println(T3)
	fmt.Println(T4)

	// 切片中append的用法
	T2 = append(T2, 1) // 向切片尾部追加内容
	fmt.Println(T2)
	T2 = append(T2, 2, 3) // 向切片尾部追加多个内容
	fmt.Println(T2)
	T3 = append(T3, T4...) // 向切片尾部追加另一个切片需要在切片名字后面加上...
	fmt.Println(T3)

	// 切片resize 重复切片
	fmt.Println(T1)
	O1 := T1[1:3]
	fmt.Println(O1)
	O1 = T1[0:2]

	// 切片的拷贝
	T5 := []int{1, 2, 3, 4, 5} // 声明一个切片并赋值
	T6 := make([]int, 10)      // 创建一个切片并声明其长度为10
	copy(T6, T5)               // 将T5切片中的内容复制到T6中
	/*
		如果T6的容量大于T5, 多出来的那些容量使用对应类型的默认值,
		如果T6的容量小于T5, 那么只复制T6容量能放下的值
	*/
	fmt.Println(T6)

	one := "hello world" // 字符串是有多个字符元素组成的, 字符串用双引号声明, 字符用单引号声明
	A1 := one[0:5]       // A1切片里每一个元素都是一个字符
	fmt.Println(A1)
	// A1[0] == 'H'		 // string类型的切片是不可以修改其元素的 但是数组可以修改(看testModifyString函数示例)
	// fmt.Println(A1)

	/*
		// 通过数组来声明切片
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
	*/
}

func testModifyString() {
	str := "hello world"
	T1 := []rune(str) // 将str转换成一个rune类型的数组, rune类型用于表示unicode字符
	T1[0] = 'H'
	fmt.Println(string(T1))
}

func main() {
	// fmt.Println("123")
	testSlice()
	testModifyString()
}
