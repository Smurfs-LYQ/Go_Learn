package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

/*
// 方法一
func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	for {
		// 使用 value, ok := <- ch 取值方式，当通道关闭时，ok = false
		res, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
*/

// 方法二
func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)

	// 使用for range循环去通道ch1中接收值
	for res := range ch1 {
		fmt.Println(res)
	}
}
