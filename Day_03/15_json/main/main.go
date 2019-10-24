package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	// 定义元信息 json tag
	/*
		`json:name` 就是元信息
		代表当json这个包来处理数据的时候，他看到的这个字段名是name，但是结构体字段首字母大写又保证了字段的可见性
	*/
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	var stu1 = Student{
		"T1",
		18,
		"男",
	}
	fmt.Printf("%T, %v\n", stu1, stu1)

	// json序列化(将stu1转换成json格式的)
	res, err := json.Marshal(stu1) // 返回一个转换后的[]byte切片和一个错误信息
	if err != nil {
		fmt.Println("转换失败: ", err)
	}
	fmt.Printf("转换后的类型: %T\n转换后的值: %v\n转换为string类型: %s\n\n", res, res, string(res))

	// 反序列化
	var stu2 Student
	json.Unmarshal([]byte(res), &stu2) // 参数1: 要转换字符串byte格式的切片  参数2: 转换之后赋值给哪个变量(这个变量的格式必须与参数1转换之后的返回值格式相对应)
	fmt.Printf("转换之后的类型: %T\n转换之后的值: %v\n", stu2, stu2)
}
