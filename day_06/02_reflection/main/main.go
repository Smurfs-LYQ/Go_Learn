package main

import (
	"fmt"
	// 反射包
	"reflect"
)

type T1 struct {
	Name string
	Age  int
}

func test(b interface{}) {
	// 获取变量的类型
	t := reflect.TypeOf(b)
	fmt.Println(t)
	// 获取变量的值
	v := reflect.ValueOf(b)
	fmt.Println(v)
	// 获取变量的类别
	fmt.Println(v.Kind())
	// 转换成interface{}类型
	iv := v.Interface()
	fmt.Printf("%T %v\n", iv, iv)
	test_1, ok := iv.(T1) // 类型断言
	if ok {
		fmt.Printf("%T %v\n", test_1, test_1)
	}
}

func project_1(b interface{}) {
	fmt.Println("type: ", reflect.TypeOf(b)) // 获取变量b的类型
	v := reflect.ValueOf(b)                  // 获取变量b的变量
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())   // 通过value中的Type方法获取变量b的类型
	fmt.Println("kind: ", v.Kind())   // 通过value中的King方法获取变量b的类别
	fmt.Println("value: ", v.Float()) // 通过value中的Float方法获取变量b的值, 但是注意：这里的Float必须是变量的数据类型

	fmt.Printf("%T %v\n", v.Interface(), v.Interface()) // 将变量转换为Interface类型
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64) // 通过类型断言将变量从Interface类型转换成float64
	fmt.Printf("%T %v\n", y, y)
}

func main() {
	// var a T1 = T1{Name: "Smurfs", Age: 21}
	// test(a)

	var a float64 = 3.4
	project_1(a)
}
