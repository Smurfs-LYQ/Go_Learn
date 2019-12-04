package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // sync.Map 自带方法: 用于设置键值对 参数:键、值
			value, _ := m.Load(key) // sync.Map 自带方法: 用于获取指定键的值 参数:键
			fmt.Printf("k=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
