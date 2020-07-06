package main

import (
	"flag"
	"fmt"
	"time"
)

// flag_demo_1 flag.Type() 示例
func flag_demo_1() {
	// 设置接收的参数信息
	name := flag.String("name", "Smurfs", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")

	// 对命令行参数进行解析
	flag.Parse()

	fmt.Println(*name, *age, *married)

	// go run main.go -name 格格巫 -age 20
}

// flag_demo_2 flag.TypeVar()
func flag_demo_2() {
	var (
		name    string
		age     int
		married bool
		delay   time.Duration
	)
	flag.StringVar(&name, "name", "Smurfs", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, delay)
}

// flag_outher flag其他函数示例
func flag_outher() {
	var (
		name string
		age  int
	)
	flag.StringVar(&name, "name", "Smurfs", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")

	// 解析命令行参数
	flag.Parse()

	fmt.Println(name, age)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

func main() {
	//flag_demo_1()

	//flag_demo_2()

	flag_outher()
}
