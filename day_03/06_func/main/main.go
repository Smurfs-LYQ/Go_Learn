package main

import "fmt"

/*
// No.1
// 自定义类型 想要成为此类型的变量要求必须与其规定的格式相同
type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Printf("%T\n", operator)
	fmt.Println(sum)
}
*/

/*
// No.2 练习
func modify(a *int) {
	*a = 100
}

func main() {
	a := 8
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}
*/

/*
// No.3 命名返回值的名字
// 创建一个函数, 接收两个int值, 返回两个值相加和相减得值
func test(a, b int) (add, del int) {
	add = a + b
	del = a - b
	return
}

func main() {
	// No.3_1 调用并接收多个返回值的函数
	// add, del := test(10, 5)
	// fmt.Println(add)
	// fmt.Println(del)

	// No.3_2 _可以用来忽略不想接收的返回值
	add, _ := test(10, 5)
	fmt.Println(add)
}
*/

// No.4 可变参数
func add(one int) int {
	fmt.Println("123")
	return "123"
}

func main() {
	fmt.Println("test")
}
