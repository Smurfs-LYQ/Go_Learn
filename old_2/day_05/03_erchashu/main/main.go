package main

import "fmt"

// 二叉树定义
/*
	type T1 struct {
		Name string
		left *T1
		right *T1
	}

	如果每个节点有两个指针分别用来指向左子树和右子树, 我们把这样的结构叫做二叉树
*/

func tranc(top *T1) {
	if top == nil {
		return
	}

	fmt.Println(top)

	tranc(top.left)
	tranc(top.right)
}

type T1 struct {
	Name  string
	left  *T1
	right *T1
}

func main() {
	var head = T1{
		Name: "head",
	}

	var left1 = T1{
		Name: "left1",
	}

	head.left = &left1

	var right1 = T1{
		Name: "right1",
	}

	head.right = &right1

	var left2 = T1{
		Name: "left2",
	}

	left1.left = &left2

	tranc(&head)
}
