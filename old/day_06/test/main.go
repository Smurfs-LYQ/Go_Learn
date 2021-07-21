package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	Name string `cn:"姓名"`
	Age  int    `cn:"年龄"`
}

func (this *T1) Set(name string, age int) {
	this.Name = name
	this.Age = age

	fmt.Println("设置成功")
}

func (this T1) Get() T1 {
	return this
}

// 反射测试
func project_1(b interface{}) {
	fmt.Println(reflect.TypeOf(b))
	this := reflect.ValueOf(b)

	switch this.Kind() {
	case reflect.Struct:
		fmt.Println("结构体")
		tye := reflect.TypeOf(b)
		// 获取字段的tag名称
		for i := 0; i < this.NumField(); i++ {
			fmt.Println(this.Field(i), "=>", tye.Field(i).Tag.Get("cn"))
		}
	case reflect.Int:
		fmt.Println("整形")
	case reflect.String:
		fmt.Println("字符串")
	}
}

func main() {
	var one T1
	// one.Set("Smurfs", 21)
	// fmt.Println(one.Get())

	// 整形测试
	project_1(123)
	// 字符串测试
	project_1("Smurfs")
	// 结构体测试
	project_1(one)

}
