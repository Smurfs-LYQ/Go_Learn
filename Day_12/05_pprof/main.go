package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func loginCode() {
	var ch chan int
	for {
		select {
		case v := <- ch:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
		}
	}
}

func main() {
	var {
		isCPUPprof bool
		isMemPprof bool
	}

	flag.BoolVar(&isCPUPprof, "cpu", false, "CPU测试")
	flag.BoolVar(&isMemPprof, "mem", false, "内存测试")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("cpu.pprof")
		if err != nil {
			fmt.Println("create cpu.pprof faield, err:", err)
			return
		}

		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0;i < 10; i++ {
		go loginCode()
	}
	time.Sleep(time.Second*20)
	if isMemPprof {
		file, err := os.Create("mem.pprof")
		if err != nil {
			fmt.Println("create mem.pprof faield, err:", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}