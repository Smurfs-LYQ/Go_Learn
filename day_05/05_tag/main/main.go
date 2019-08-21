package main

import (
	"encoding/json"
	"fmt"
)

/*
// 我们可以为struct中的每个字段，写上一个tag。这个tag可以通过反射的机制获取到，最常用的场景就是json序列化和反序列化

type T1 struct {
	Name string "this is name field" // 数据类型后面双引号部分的内容就是这个字段的tag
	Age  int    "this is age field"
}
*/

type T1 struct {
	Name string `json:"name"` // 这里这个tag的意思是：这个元素在json进行打包操作的时候元素名字为小写的name
	Age  int    `json:"age"`
}

func main() {
	one := T1{
		Name: "test",
		Age:  1,
	}

	// fmt.Println(one)
	// 进行json打包操作
	data, err := json.Marshal(one) // 它会返回两个内容，分别为打包之后的内容和错误信息，而打包之后的内容是 []byte 类型的数组
	if err != nil {
		fmt.Println("错误异常：", err)
		return
	}

	fmt.Println(string(data))
}
