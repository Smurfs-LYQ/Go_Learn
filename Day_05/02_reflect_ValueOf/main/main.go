package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x) // 获取接口的值信息
	k := v.Kind()           // 拿到值对应的种类
	fmt.Printf("%#v, %v\n", v, k)
	switch k {
	case reflect.Int:
		// v.Int() 从反射中获取整型的原始值，然后通过int()强制类型转换
		fmt.Printf("这是一个int类型的值，值为: %d\n", int(v.Int()))
	case reflect.String:
		// v.String() 从反射中获取字符串的原始值，然后通过string()强制类型转换
		fmt.Printf("这是一个string类型的值，值为: %s\n", string(v.String()))
	case reflect.Bool:
		// v.Bool() 从反射中获取布尔值的原始值，然后通过bool()强制类型转换
		fmt.Printf("这是一个bool类型的值，值为: %t\n", bool(v.Bool()))
	default:
		fmt.Println(k)
	}
}

func main() {
	reflectValue(1234)
	reflectValue("1234")
	reflectValue(false)
	reflectValue(12.34)

	// 将int类型的原始值转换为revlect.Value类型
	fmt.Printf("type : %T\n", reflect.ValueOf(10))
}
