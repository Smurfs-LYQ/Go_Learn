package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)                                         // 可以拿到x的动态类型信息
	fmt.Printf("type:%v name:%v kind:%v\n", t, t.Name(), t.Kind()) // 原理就是用的反射，代码补全的原理也是反射
}

// Cat 定义一个猫的结构体
type Cat struct {
	name string
}

// People 定义一个人的结构体
type People struct {
	name string
	age  int
}

func main() {
	/*
		reflectType(123)
		reflectType(false)
		reflectType("hello world")
		reflectType([3]int{1, 2, 3})
		reflectType(map[int]string{})
	*/

	var a *float32
	reflectType(a)
	Tom := Cat{"tom"}
	reflectType(Tom)
	Man := People{"Smurfs的格格巫", 18}
	reflectType(Man)
}
