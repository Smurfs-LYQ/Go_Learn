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

func (s T1) Set(name string, age int) {
	s.Name = name
	s.Age = age

	fmt.Println("设置完成")
}

func (s T1) Get() {
	fmt.Println("-----Start-----")
	fmt.Println(s)
	fmt.Println("------End------")
}

func test(b interface{}) {
	// 获取变量的类型
	t := reflect.TypeOf(b)
	fmt.Println(t)
	// 获取变量的值，并将变量转换为reflect.value类型
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
	v := reflect.ValueOf(b)                  // 获取变量b的变量, 并将变量转换为reflect.value类型
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())   // 通过value中的Type方法获取变量b的类型
	fmt.Println("kind: ", v.Kind())   // 通过value中的King方法获取变量b的类别
	fmt.Println("value: ", v.Float()) // 通过value中的Float方法获取变量b的值, 但是注意：这里的Float必须是变量的数据类型

	fmt.Printf("%T %v\n", v.Interface(), v.Interface()) // 将变量转换为Interface类型
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64) // 通过类型断言将变量从Interface类型转换成float64
	fmt.Printf("%T %v\n", y, y)
}

func project_2(b interface{}) {
	// 获取变量的值，并将变量转换为reflect.value类型
	this := reflect.ValueOf(b)
	// 通过反射更改变量的值 SetInt() SetString()...
	this.Elem().SetInt(100) // Elem()函数返回this持有的接口保管的值得value封装，或者this持有的指针指向的值的value封装(如果是指针就类似于*this获取指针指向的变量得操作)。
	fmt.Println(this.Kind(), this)
}

func project_3(b interface{}) {
	// 获取变量的值，并将变量转换为reflect.value类型
	this := reflect.ValueOf(b)

	// 注意case后面跟的值前面的reflect是因为变量被reflect.ValueOf转换成了reflect包中的类型
	switch this.Kind() {
	case reflect.Struct:
		// 通过反射操作结构体
		// 获取结构体字段的数量
		fmt.Println("元素的数量: ", this.NumField())

		// 获取结构体方法的数量
		fmt.Println("方法的数量: ", this.NumMethod())

		// 调用结构体中的第二个方法,
		var params []reflect.Value
		this.Method(0).Call(params) // 这里的1代表的是结构体中的第二个方法, 0代表第一个方法, 以此类推
	case reflect.Int:
		fmt.Println("type is Int")
	default:
		fmt.Println("other")
	}
}

func main() {
	// var a T1 = T1{Name: "Smurfs", Age: 21}
	// test(a)

	// var a float64 = 3.4
	// project_1(a)

	// var a int = 1
	// fmt.Println(a)
	// project_2(&a)
	// fmt.Println(a)

	var a T1 = T1{
		Name: "Smurfs",
		Age:  21}
	project_3(a)
}
