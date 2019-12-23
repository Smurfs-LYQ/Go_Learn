package main

import "time"

func main() {
	for i := 0; i <= 100; i++ {
		//go并发执行
		go goroute(i)
	}

	time.Sleep(time.Second) //time.Second是time包的默认时间单位(秒)
}
