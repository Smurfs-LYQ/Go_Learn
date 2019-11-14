package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		var T1 string = "Hello"
		Test_1 := reflect.ValueOf(&T1)
		fmt.Println(Test_1.IsNil())

		var T2 string = "world"
		Test_2 := reflect.ValueOf(T2)
		fmt.Println(Test_2.IsValid())
	*/

	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员: ", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法: ", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// c := map[string]int{"Smurfs": 1}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键: ", reflect.ValueOf(c).MapIndex(reflect.ValueOf("Smurfs")).IsValid())
}
