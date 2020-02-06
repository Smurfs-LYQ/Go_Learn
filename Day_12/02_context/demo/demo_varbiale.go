package demo

import "sync"

import "fmt"

import "time"

var wg_1 sync.WaitGroup

// 声明一个全局变量
var status bool

func do_one() {
	defer wg_1.Done()
	for {
		fmt.Println("waiting...")
		time.Sleep(time.Second)
		if status {
			break
		}
	}
}

// Variable_demo 使用变量退出goroutine实例
func Variable_demo() {
	wg_1.Add(1)
	go do_one()

	time.Sleep(time.Second * 5)
	status = true
	wg_1.Wait()
	fmt.Println("over")
}
