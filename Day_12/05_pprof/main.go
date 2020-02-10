// runtime_pprof/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:

		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

/*
运行: go tool pprof pprof文件名

运行结果:
	```
	Type: cpu
	Time: Feb 10, 2020 at 2:08pm (CST)
	Duration: 20.23s, Total samples = 59.96s (296.37%)
	Entering interactive mode (type "help" for commands, "o" for options)
	(pprof)
	```

这是一个可交互的命令行，可以通过"top 数字"形式来查看消耗系统资源的前几名程序

例如输入: top 3
	```
	Showing nodes accounting for 44.14s, 91.79% of 48.09s total
	Dropped 15 nodes (cum <= 0.24s)
	Showing top 3 nodes out of 4
		flat  flat%   sum%        cum   cum%
		18.06s 37.55% 37.55%     36.76s 76.44%  runtime.selectnbrecv
		14.90s 30.98% 68.54%     16.17s 33.62%  runtime.chanrecv
		11.18s 23.25% 91.79%     47.94s 99.69%  main.logicCode
	```

	其中:
		flat：当前函数占用CPU的耗时
		flat：当前函数占用CPU的耗时百分比
		sun%：函数占用CPU的耗时累计百分比
		cum： 当前函数加上调用当前函数的函数占用CPU的总耗时
		cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
		最后一列：函数名称

还可以使用"list 函数名"命令查看具体的函数分析，例如执行"list loginCode"查看我们编写的函数的详细分析

	```
	Total: 48.09s
	ROUTINE ======================== main.logicCode in /Users/smurfs/Golang/src/Go_Learn/Day_12/05_pprof/main.go
		11.18s     47.94s (flat, cum) 99.69% of Total
			.          .     12:// 一段有问题的代码
			.          .     13:func logicCode() {
			.          .     14:   var c chan int
			.          .     15:   for {
			.          .     16:           select {
		11.18s     47.94s     17:           case v := <-c:
			.          .     18:                   fmt.Printf("recv from chan, value:%v\n", v)
			.          .     19:           default:
			.          .     20:
			.          .     21:           }
			.          .     22:   }
	```

	通过分析发现大部分CPU资源被17行占用，我们分析出select语句中的default没有内容会导致上面的"case v:=<-c:"一直执行。我们在default分支添加一行"time.Sleep(time.Second)"即可。

退出: exit
*/
