package main

import "fmt"

/*
	1. map简介: key-value的数据结构, 又叫字典或关联数组
	2. 声明:
		var name map[keytype]valuetype
		例:
			var a map[string]string				// key为字符串, value为字符串
			var a map[string]int				// key为字符串, value为整形
			var a map[int]string				// key为整形, value为字符串
			var a map[string]map[string]string 	// key为字符串, value为另一个map
		注意: 声明不会分配内存, 初始化需要make
*/

func testMap_1() {
	/*
		// 声明map类型
		var a map[int]int
		// 初始化
		a = make(map[int]int, 10) // 初始化为map[string]string类型, 长度为10
	*/

	// 简短写法
	a := make(map[int]int, 10)

	// 赋值
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	// a[5] = 6
	// a[6] = 7
	// a[7] = 8
	// a[8] = 9
	// a[9] = 10

	fmt.Println(a)
}

func testMap_2() {
	// 声明并初始化赋值一个map
	// var a map[int]string = map[int]string{
	a := map[int]string{
		1: "one", // 注意这里的逗号必须写
		2: "two",
		3: "thr",
	}
	fmt.Println(a)
}

func testMap_3() {
	/*
		// 声明并初始化一个map
		var a = make(map[int]map[int]string, 100)
		// 给map a的子map赋值需要对其重新初始化
		a[0] = make(map[int]string)
		a[1] = make(map[int]string)
		// 赋值
		a[0][0] = "0_one"
		a[0][1] = "0_two"
		a[1][0] = "1_one"
		fmt.Println(a)
	*/

	// 声明并初始化一个map
	a := make(map[int]map[int]string, 10)

	//给a map的子map初始化赋值
	a[0] = map[int]string{
		0: "0_One",
		1: "0_Two",
	}

	a[1] = map[int]string{
		0: "1_One",
	}

	fmt.Println(a)
}

func testMapFind() {
	a := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	/*
		查找
			格式: 值变量名, 查找结果变量名 := map变量名[键]
			返回值:
				如果查到了, 返回对应的value和true
				如果查不到, 返回 空字符串 和false
	*/
	val, ok := a[3]
	fmt.Println(val, ok)
}

func testMapRange() {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "thr",
	}

	for key, val := range a {
		fmt.Printf("key= %d\nval= %s\n\n", key, val)
	}
}

func testMapDelete() {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "thr",
	}

	fmt.Println(a, "删除前长度为:", len(a))
	// 删除指定map的指定元素
	delete(a, 3) //map变量名, 要删除的键
	fmt.Println(a, "删除后长度为:", len(a))
}

func testMapReversal() {
	// key val反转
	a := map[string]int{
		"one": 1,
		"two": 2,
		"thr": 3,
	}
	fmt.Println(a)

	// 重新创建并初始化一个map用来存储反转后的map
	b := make(map[int]string)

	for key, val := range a {
		b[val] = key
	}

	fmt.Println(b)
}

func main() {
	// testMap_1()
	// testMap_2()
	// testMap_3()

	// testMapFind()
	// testMapRange()
	// testMapDelete()

	testMapReversal()
}
