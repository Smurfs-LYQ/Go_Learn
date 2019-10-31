package main

import (
	"fmt"
	"math/rand"
	"time"
)

// writer 定义一个接口
type writer interface {
	Write([]byte) string
}

// Student 定义实现接口的结构体
type Student struct {
	name string
}

// Student结构体实现接口中指定的方法
func (s Student) Write(one []byte) string {
	return fmt.Sprintf("%s 正在写: %s", s.name, string(one))
}

// Teacher 声明一个老师的结构体
type Teacher struct {
	name string
}

// Teacher结构体实现接口中指定的方法
func (t Teacher) Write(one []byte) string {
	return fmt.Sprintf("%s 正在写: %s", t.name, string(one))
}

func main() {
	// 为了使随机数每次都产生不一样的数字，所以需要用Seed函数初始化一下
	rand.Seed(time.Now().Unix())

	/*
		// 演示一
		// 实例化接口
		var student writer
		fmt.Printf("%P %T\n", &student, student)

		// 初始化结构体
		var stu1 = Student{
			name: "Smrufs",
		}
		fmt.Printf("%P %T\n", &stu1, stu1)

		// 因为Student这个接口体实现了writer接口的规范，所以Student实例化的对象可以赋值给writer实例化的变量
		student = stu1
		fmt.Printf("%P %T\n", &student, student)
		student.Write([]byte("格格巫"))
	*/

	// 演示二
	// 实例化接口
	var list []writer

	// 初始化结构体
	var stu = Student{
		name: "Smurfs",
	}
	// 初始化结构体
	var teacher = Teacher{
		name: "体育老师",
	}

	// 将不同结构体放入到接口类型的切片中(结构体必须实现了接口指定的规范)
	list = append(list, stu, teacher)

	// 通过接口调用方法
	for _, v := range list {
		str := []byte(fmt.Sprintf("%d", rand.Intn(10000)))
		res := v.Write(str)
		fmt.Println(res)
	}
}
