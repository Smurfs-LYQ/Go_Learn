package main

import (
	"errors"
	"fmt"
)

// 自己抛出异常错误
func initConfig() (err error) {
	return errors.New("初始化config配置文件出错") // 返回一个错误需要调用errors包中的New变量
}

// panic和recover测试用
func test_painc() {
	// 这是一个有问题的函数
	defer func() {
		if err := recover(); err != nil {
			// fmt.Println("问题: 整数被除零")
			println(err)
		}
	}()

	// b := 0
	// a := 100 / b
	// fmt.Println(a)

	err := initConfig()
	if err != nil {
		panic(err)
	}

	return
}

// Go内置函数
func main() {
	var one string

	// close 主要用来关闭channel
	// one_1

	// len 用来求长度，比如：string、array、slice、map、channel
	one = "hello golang"
	fmt.Println(len(one))

	// new 用来分配内存，主要用来分配值类型，比如：int、struct。返回的是指针地址。也可以称为初始化一个值类型的变量
	fmt.Println(one)
	one_2 := new(string) // 参数：变量的类型
	*one_2 = "hello golang"
	fmt.Println(*one_2)

	// make 用来分配内存，主要用来分配引用类型，比如chan、map、slice。也可以称为初始化一个引用类型的变量
	one_3 := make([]int, 5) // 参数：变量的类型，变量的容量(如果容量没有被使用默认都为对应类型的空值，比如int为0)
	// one_3 = []int{1, 2, 3}
	fmt.Println(one_3)
	/* 注: make如果不填写长度和容量依然可以使用, 使用的是其默认值,
	但是如果长度或这容量不够用的时候回自动进行扩容, 比如默认内存为5,
	扩容就是5+5+5..., 如果明确知道自己需要的容量是多少, 那就设置上
	默认的容量, 这样性能回更高, 省区了容量不足时扩容的操作
	*/

	// append 用来追加元素到 数组、slice中
	one_4 := []int{1, 2, 3, 4, 5}
	fmt.Println(one_4)
	one_4 = append(one_4, 6, 7, 8, 9, 10)
	fmt.Println(one_4)

	// 将一个slice中的元素追加到指定slice中
	test_1 := []string{"hello", "golang"}
	fmt.Println(test_1)
	test_2 := []string{"hello", "world"}
	test_1 = append(test_1, test_2...) // 这里的 test_2... 代表的就是将test_2中的所有元素
	fmt.Println(test_1)

	// panic和recover 用来做错误处理
	// 虽然如果代码出现错误程序也会报错，但是程序自己的报错是直接提示错误并且关闭程序的，但是自己捕获错误的话就可以避免程序被关闭的问题
	// one_5
	test_painc()
}
