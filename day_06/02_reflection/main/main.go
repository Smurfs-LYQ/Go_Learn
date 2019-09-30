package main

import (
	"fmt"
	// 反射包
	"reflect"
)

type T1 struct {
	Name string `json:"name"`
	Age  int
}

func (s *T1) Set(name string, age int) (string, int) {
	s.Name = name
	s.Age = age

	return s.Name, s.Age
}

func (s *T1) Get() *T1 {
	return s
}

func (s T1) Test_T1() {
	fmt.Println("测试方法Test_T1")
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
	fmt.Println(this)

	// 注意case后面跟的值前面的reflect是因为变量被reflect.ValueOf转换成了reflect包中的类型
	switch this.Elem().Kind() {
	case reflect.Struct:
		// 通过反射操作结构体
		// 获取结构体字段的数量
		fmt.Println("元素的数量: ", this.Elem().NumField()) // 这里加Elem的原因是因为this是一个"结构体指针", 通过Elem获取到"结构体指针"的"结构体", 然后通过NumField获取字段数量, 如果this不是"结构体指针"就不需要Elem了

		// 获取结构体方法的数量
		fmt.Println("方法的数量: ", this.NumMethod()) // 如果this为"结构体指针"那么他可以获取到此结构体所有首字母大写的方法, 如果this为"结构体"那么他就只能获取到所有首字母大写的非指针型方法

		// 调用结构体中的字段
		for i := 0; i < this.Elem().NumField(); i++ {
			fmt.Println(i, this.Elem().Field(i))
		}

		// 修改结构体中字段的值
		this.Elem().Field(0).SetString("修改测试") // 还有SetInt等...

		// 调用结构体中的第一个方法, 带参数
		setNameMethod_one := this.MethodByName("Set")                                   // 设置调用方法的名字
		args_one := []reflect.Value{reflect.ValueOf("Smurfs的格格巫"), reflect.ValueOf(21)} // 创建参数列表
		res_one := setNameMethod_one.Call(args_one)                                     // 调用方法, 并传入参数, 注意Call方法的返回值是一个"reflect.Value类型的slice"
		fmt.Printf("%v %T\n", res_one[0], res_one)

		// 调用结构体中的第二个方法, 不带参数
		setNameMethod_two := this.MethodByName("Get")
		args_two := make([]reflect.Value, 0)
		res_two := setNameMethod_two.Call(args_two)
		fmt.Println(res_two[0])
	case reflect.Int:
		fmt.Println("type is Int")
	default:
		fmt.Println("other")
	}
}

func project_4(b interface{}) {
	tye := reflect.TypeOf(b)

	// json包中的Marshal方法就是通过反射实现的 // 如果这个json这个地方不明白去看一下day_05/05_tag
	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Println(tag)
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
	//project_3(&a)
	project_4(&a)
}
