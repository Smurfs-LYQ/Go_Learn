package main

import "fmt"

// Project "Go_Learn/day_03/06_func/project"

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
/*
// No.4_1 0-无限个参数
func add(one ...int) int {
	sum := 0
	for _, v := range one {
		sum += v
	}
	return sum
}
*/

/*
// No.4_2 最少一个参数
func add(one int, two ...int) int {
	var sum int = one
	for i := 0; i < len(two); i++ {
		sum += two[i]
	}
	return sum
}

// No.4_2 最少两个参数
func add(one, two int, thr ...int) int {
	var sum int = one
	for i := 0; i < len(thr); i++ {
		sum += thr[i]
	}
	return sum + two
}

func main() {
	// No.4_1
	// fmt.Println(add(10, 1, 2, 3, 4, 5))
	// No.4_2
	fmt.Println(add(10, 10, 1, 2, 3, 4, 5))
}
*/

// No.5_1 defer使用示例
func defer_test() {
	i := 0
	// 它返回的结果是0, 因为它在定义时, i的值是0, 他会在程序的最后执行
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

// No.6 匿名函数
func test(a, b int) {
	res := func(a1, b1 int) int { // 匿名函数不需要写函数名
		return (a1 + b1)
	}(a, b) // 匿名函数直接在函数的结尾就可以直接加 () 进行调用, 也可以通过变量名"res"进行调用
	fmt.Println(res)
}

func main() {
	// Project_1
	// fmt.Println(project.Project_1("h", "e", "l", "l", "o"))

	// No.5_1 defer的使用示例
	// defer_test()

	// No.6 匿名函数
	test(1, 1)
}
