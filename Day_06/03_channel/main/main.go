package main

import "fmt"

// T1 channl练习
func T1() {
	// 创建管道
	var ch1 chan int
	// 初始化管道
	ch1 = make(chan int, 3)
	// 向管道内传递值
	ch1 <- 299
	ch1 <- 199
	ch1 <- 99
	// 获取并打印出管道中的值，然后直接丢弃掉
	fmt.Printf("%#v\n", <-ch1)
	fmt.Printf("%#v\n", <-ch1)
	fmt.Printf("%#v\n", <-ch1)
	// 关闭通道
	close(ch1)

	// 创建并初始化一个管道
	ch2 := make(chan string, 5)
	// 向管道内传递至
	ch2 <- "Smurfs"
	ch2 <- "Smurfs的格格巫"
	// 从ch2中接收值，保存变量到res1变量中
	res1 := <-ch2
	fmt.Printf("%#v\n", res1)
	// 关闭管道
	close(ch2)
	// 关闭管道后再次接收管道中的值, 并保存到res2变量中
	res2 := <-ch2
	fmt.Printf("%#v\n", res2)
	fmt.Printf("%#v\n", <-ch2) //管道关闭之后依然可以接收值，直到管道中的值都被接收完，之后再接受到就是对于类型的零值
}

// channel

func main() {
	// 定义一个ch1变量，是一个channel类型，这个channel内部传递的数据是int类型
	var ch1 chan int

	// 定义一个ch2变量，是一个channel类型，这个channel内部传递的数据是string类型
	var ch2 chan string

	// channel 是引用类型
	fmt.Println(ch1, ch2)
	// make函数初始化(分配内存): slice map channel
	ch3 := make(chan int, 1) // 容量为1的管道，最多只能放一个int类型的数据

	// 通道的操作: 发送、接收、关闭
	// 发送和接收都用一个符号:	<-
	ch3 <- 10 // 把10发送到ch3中
	// <-ch3        // 从ch3中接收值，直接丢弃
	res := <-ch3 // 从ch3中接收值，保存到变量res中
	fmt.Println(res)
	// 关闭
	/*
		1. 关闭的通道再接收, 能取到对应类型的零值
		2. 向关闭的通道中发送值, 会引发panic
		3. 关闭一个已经关闭的通道, 会引发panic
	*/
	close(ch3)
}
