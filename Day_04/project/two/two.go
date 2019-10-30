package two

import (
	"fmt"
	"time"
)

// Two_1 编写程序统一一段代码的执行耗时时间，单位精确到微妙
func Two_1() {
	start := time.Now().UnixNano() / 1000
	time.Sleep(1234 * time.Microsecond)
	stop := time.Now().UnixNano() / 1000
	fmt.Println(stop - start)
}

// Two_2 编写程序统一一段代码的执行耗时时间，单位精确到微妙，使用time包自带的方法
func Two_2() {
	now := time.Now()
	time.Sleep(1234 * time.Microsecond)
	fmt.Println(time.Since(now)) // 使用time包自带的Since方法计算从传入的时间到程序执行到现在耗费了多少时间
}
