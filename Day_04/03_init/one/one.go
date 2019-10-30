package one

import "fmt"

var T1 = "在init函数触发之前声明一个字符串\"变量\""

const T2 = "在init函数触发之前声明一个字符串\"常量\""

func init() {
	fmt.Printf("%s\n%s\n", T1, T2)
	fmt.Println("恭喜成功导入了one包")
}

func One() {
	fmt.Println("One输出")
}
