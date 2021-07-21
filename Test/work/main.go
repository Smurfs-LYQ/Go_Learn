package main

import (
	"Go_Learn/Test/work/work"
	"fmt"
	"log"
	"sync"
)

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	// time.Sleep(time.Second)
}

func main() {
	// 使用两个goroutine来创建工作池
	p := work.New(5)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行。当Run返回的时候，我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
		fmt.Printf("########## %d ###########", i)
	}

	wg.Wait()
	log.Println("########## die ###########")

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
