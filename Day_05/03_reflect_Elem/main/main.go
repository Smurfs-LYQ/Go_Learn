package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值
func reflectElem(x interface{}) {
	// 首先将x转换为reflect.Value类型的变量v
	v := reflect.ValueOf(x)

	// 通过reflect包中的Kind方法获取变量v的类型
	if v.Kind() == reflect.Ptr {
		/*
			修改变量v的值
				因为函数传参都是值传递，所以如果想修改其本身的话需要将其对应的内存地址传入
				但是函数接收的是一个interface类型，所以通过将其转换为reflect.Value类型
				然后再调用reflect的Elem函数获取他内存指向的值，之后再调用reflect中设置
				值得函数就可以了(例如: SetInt SetString...)
		*/
		v.Elem().SetInt(123)
	}
}

func main() {
	var a int64 = 100
	reflectElem(&a)
	fmt.Println(a)
}
