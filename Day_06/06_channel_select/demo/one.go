package demo

import (
	"fmt"
	"time"
)

var ch1 = make(chan string, 3)
var ch2 = make(chan string, 3)

func T1(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("T1: %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func T2(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("T2: %d\n", i)
		time.Sleep(time.Second * 2)
	}
}

func One() {
	go T1(ch1) // 往ch1这个通道中放T1开头的字符串
	go T2(ch2) // 往ch2这个通道中放T1开头的字符串

	for {
		select {
		case res := <-ch1:
			fmt.Println(res)
		case res := <-ch2:
			fmt.Println(res)
		default:
			fmt.Println("暂时取不到值")
		}
		time.Sleep(time.Second)
	}
}
