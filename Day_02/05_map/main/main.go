package main

import "fmt"

func main() {
	// 声明一个map类型，但不初始化
	var T1 map[int]string
	fmt.Printf("T1 == nil: %t\n", T1 == nil)

	// 声明一个map的同时初始化
	var T2 = map[int]string{
		1: "one",
		2: "two",
		3: "thr",
	}
	fmt.Println(T2)

	// 声明一个map类型，并使用make初始化
	var T3 map[string]int
	T3 = make(map[string]int, 3) // 第一位为初始化的类型，第二位为容量

	// map中如何添加键值对
	T3["one"] = 1
	T3["two"] = 2
	T3["thr"] = 3
	fmt.Println(T3)

	// 判断指定key的在不在对应的map中
	v, ok := T3["one"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有")
	}
	/*
		如果指定的key在对应的map中
			v = 1 (指定的key所对应的值)
			ok = true
		如果指定的key不在对应的map中
			v = 对应map的值类型的零值
			ok = false
	*/

	// 遍历map中的key和value
	for k, v := range T3 {
		fmt.Println(k, v)
	}

	// 删除map中指定的键值对
	delete(T3, "thr")
	fmt.Println(T3)
}
