package main

import (
	"fmt"
)

type animal interface {
	speak()
	move()
}

// 定义一个猫的结构体
type cat struct {
	name string
}

// 定义使用值接收者的方法
func (c cat) speak() {
	fmt.Println("喵喵喵")
}

// 定义使用值接收者的方法
func (c cat) move() {
	fmt.Println("猫步前进")
}

// 定义一个狗的结构体
type dog struct {
	name string
}

// 定义使用指针接收者的方法
func (d *dog) speak() {
	fmt.Println("汪汪汪")
}

// 定义使用指针接收者的方法
func (d *dog) move() {
	fmt.Println("飞速移动")
}

func main() {
	var x animal

	// 值接收者创建的方法
	var hh = cat{"花花"}
	x = hh
	var tom = &cat{"汤姆猫"}
	x = tom

	x.speak()
	x.move()
	tom.speak() // 语法糖 (*tom).speak()
	tom.move()  // 语法糖 (*tom).move()

	fmt.Println()

	// 指针接受者创建的方法
	/*
		var peter = dog{"皮特"} // peter是dog类型
		x = peter             // x不可以接收dog类型
	*/
	var peter = &dog{"皮特"} // peter是*dog类型
	x = peter              // x可以接收*dog类型

	x.speak()
	x.move()
}
