package main

import "fmt"

/*
接口：
	1. interface类型默认是一个指针
	2. interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量
	3. Golang中的接口，不需要显性的实现。只要一个变量，含有接口类型中的“所有方法”，那么这个变量就实现了这个接口。因此，Golang中没有implement类似的关键字
	4. 如果一个变量含有了多个interface类型(接口)的方法，那么这个变量就实现了多个接口
	5. 如果一个变量只含有一个interface的部分方法, 那么这个变量就没有实现这个接口
定义:
	type 接口名 interface {
		方法名1(参数列表) 返回值列表
		方法名2(参数列表) 返回值列表
	}
多态：
	字面意思即“多种形态”。一个事物的多种形态，都可以按照统一的接口进行操作即为多态。在面向对象语言中，接口的多种实现方式即为多态。
	车 可以代表汽车和火车等，所以 车 这个概念就是典型的多态，因为他不仅可以代表汽车 也可以代表火车等
	例如我们下面定义了一个T1的接口，其中有两个方法。后面又分别定义了Man和WoMan的结构体，并且他们都实现了T1接口中定义的两个方法。那么接口T1既可以通过Man结构体来实现自己，也可以通过WoMan结构体来实现自己，这种有多种实现方式的接口即为多态
*/

// 定义一个接口
type T1 interface {
	test()
	sleep()
}

// 定义一个结构体
type Man struct {
	Name string
	Age  int
}

type WoMan struct {
	Name string
	Age  int
}

// 实现接口
func (obj Man) test() {
	fmt.Println("调用了test接口")
	fmt.Println(obj.Name)
	fmt.Println(obj.Age)
}

func (obj Man) sleep() {
	fmt.Println("男人睡觉")
}

func (obj WoMan) test() {
	fmt.Println("调用了test接口")
	fmt.Println(obj.Name)
	fmt.Println(obj.Age)
}

func (obj WoMan) sleep() {
	fmt.Println("女人睡觉")
}

func main() {
	var Test_1 T1

	var one Man = Man{
		Name: "Smurfs",
		Age:  21,
	}

	var two WoMan = WoMan{
		Name: "XXX",
		Age:  21,
	}

	// one对应的结构体实现了test接口
	Test_1 = one // 这里的赋值操作就相当于把one.test方法赋值给了Test_1.test接口
	// 调用了test接口
	Test_1.test()
	Test_1.sleep()

	fmt.Println()

	Test_1 = two
	Test_1.test()
	Test_1.sleep()
}
