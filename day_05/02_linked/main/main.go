package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

/*
链表定义:
	type T1 struct {
		Name string
		Next *T1 // 下一个节点的地址
	}

	1. 每个节点包含下一个节点的地址, 这样把所有的节点串起来了, 通常把链表中的第一个节点叫做链表头
	2. 分为 单列表 和 双列表
*/

// 创建结构体
type T1 struct {
	Name string // 元素1 姓名
	Age  int    // 元素2 年龄
	next *T1    // 元素3 下一个节点的内存地址
}

// 创建循环打印链表的方法
func trans(head *T1) { // 形参意义为一个可供for循环持续循环的载体, 值为链表头的内存地址
	for head != nil { // 判断载体后面是否还有节点
		fmt.Println(*head)  // 打印当前节点
		head = (*head).next // 修改载体的值, 值为next元素指向的下一个节点的内存地址
	}
}

// 尾部插入
func install_end(end *T1) {
	fmt.Println("尾部插入节点测试")
	for i := 0; i < 5; i++ {
		var test = T1{
			Name: fmt.Sprintf("end_%d", i),
			Age:  rand.Intn(100),
		}

		end.next = &test
		end = &test
	}
}

// 头部插入
func install_top(top *T1) *T1 {
	for i := 0; i < 5; i++ {
		var test = T1{
			Name: fmt.Sprintf("top_%d", i),
			Age:  rand.Intn(100),
		}

		test.next = top
		top = &test
	}
	return top
}

func main() {
	// 创建节点 one
	var one = T1{
		Name: "Test_One",
		Age:  18,
	}

	// 创建节点 two
	var two = T1{
		Name: "Test_Two",
		Age:  19,
	}

	// 创建节点 thr
	var thr = T1{
		Name: "Test_Thr",
		Age:  20,
	}

	// 给头结点的next元素赋值
	one.next = &two
	two.next = &thr

	// 调用打印函数
	trans(&one)

	fmt.Println()

	// 尾部插入节点测试
	install_end(&thr)
	// 调用打印函数
	trans(&one)
	fmt.Println()

	// 头部插入节点测试
	fmt.Println("头部插入节点测试")
	top := install_top(&one)
	// 调用打印函数
	trans(top)
}
