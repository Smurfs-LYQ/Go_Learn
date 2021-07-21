package main

import (
	"fmt"
)

// 定义一个学生接口
type Student interface {
	Write(string) string
}

// 定义一个老师接口
type Teacher interface {
	Speak(string) string
}

// 定义一个人的接口，并嵌套学生和老师接口
type Man interface {
	Student
	Teacher
}

// 定义一个结构体
type One struct {
	name string
}

// 结构体实现学生接口中的Write方法
func (o One) Write(str string) (res string) {
	res = fmt.Sprintf("%s 正在写 %s", o.name, str)
	return
}

// 结构体实现老师接口中的Speak方法
func (o One) Speak(str string) string {
	res := fmt.Sprintf("%s 正在讲 %s", o.name, str)
	return res
}

func main() {
	// 实例化Man接口
	var Smurfs Man

	// 初始化One结构体
	var S1 = One{
		name: "某个男学生",
	}
	// 初始化One结构体
	var T1 = One{
		name: "体育老师",
	}

	/*
		因为Man这个接口嵌套了Student和Teacher这两个接口，所以只要One结构体同时实现了Student和Teacher这两个接口，One也就实现了Man这个接口
		因为One实现了Man接口，所以他的对象S1和T1都可以赋值给接口，通过接口调取想调用的方法
	*/
	Smurfs = S1
	fmt.Println(Smurfs.Write("某个女学生的作业"))
	Smurfs = T1
	fmt.Println(Smurfs.Speak("数学课"))

}
