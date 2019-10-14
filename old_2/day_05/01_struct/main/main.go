package main

import "fmt"

/*
struct结构体
	1. 用来自定义复杂数据结构
	2. struct里面可以包含多个字段(属性)
	3. struct类型可以定义方法, 注意和函数的区分
	4. struct类型是值类型
	5. struct类型可以嵌套
	6. Go语言没有class类型, 只有struct类型
	7. struct中的元素名, 如果开头大写 外部调用结构体的时候可以访问此元素, 反之亦然
	8. struct的内存布局: struct中的所有字段在内存中是连续的
*/

// struct声明
type T1 struct {
	Name string
	Age  int
}

type Test_1 T1 // 给 结构体 或 数据类型 设置别名

func main() {
	// struct定义的三种形式
	// var one T1
	// var one *T1 = new(T1) // 因为这里创建的是指针类型的, 需要手动分配内存空间
	// var one *T1 = &T1{}	 // 因为这里创建的是指针类型的, 需要手动分配内存空间
	// 其中第二种和第三种返回的都是指向结构体的指针, 访问形式如下
	// one.Name
	// one.Age
	// 或者
	// (*one).Name
	// (*one).Age

	// 定义结构体_1
	var one T1
	// 调用结构体
	one.Name = "Test_One"
	one.Age = 18
	fmt.Println(one)
	fmt.Printf("Name: %p\n", &one.Name)
	fmt.Printf("Age: %p\n", &one.Age)

	// 定义结构体_2
	// var two T1 = T1{
	// var two = T1{
	var two *T1 = &T1{
		Name: "Test_Two",
		Age:  19,
	}
	fmt.Println(*two)
}
