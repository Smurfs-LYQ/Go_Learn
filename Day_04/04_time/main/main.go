package main

import (
	"fmt"
	"time"
)

func time1() {
	fmt.Println("################时间类型################")
	now := time.Now() // 获取当前时间
	fmt.Printf("%T\n%v\n", now, now)

	fmt.Println(now.Year())   // 年
	fmt.Println(now.Month())  // 月
	fmt.Println(now.Day())    // 日
	fmt.Println(now.Hour())   // 小时
	fmt.Println(now.Minute()) // 分钟
	fmt.Println(now.Second()) // 秒
}

func time2() {
	fmt.Println("################时间戳################")
	now := time.Now()           // 获取当前时间
	fmt.Println(now.Unix())     // 时间戳
	fmt.Println(now.UnixNano()) // 纳秒时间戳
}

func time3() {
	fmt.Println("################将时间戳转为时间格式################")
	now := time.Now()                   // 获取当前时间
	UnixTimestamp := now.Unix()         // 获取当前时间戳
	time := time.Unix(UnixTimestamp, 0) // 将当前时间戳转换为时间格式
	fmt.Println(time)
	fmt.Println(time.Year())   // 年
	fmt.Println(time.Month())  // 月
	fmt.Println(time.Day())    // 日
	fmt.Println(time.Hour())   // 时
	fmt.Println(time.Minute()) // 分
	fmt.Println(time.Second()) // 秒

	fmt.Printf("%4d-%2d-%0d %02d:%02d:%02d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

func time4() {
	fmt.Println("################定时器################")
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}

func time5() {
	fmt.Println("################时间格式化################")
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	fmt.Println(now.Format("2006-01-02 15:04:05.000")) // 05表示秒 .000表示毫秒

	// 格式化为12小时制格式
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
}

func time6() {
	// 定义一个时间对象
	now := time.Now()
	// 调用time对象的Add方法，获得当前时间加1小时后的时间
	later := now.Add(time.Hour)
	fmt.Println(now)
	fmt.Println(later)
}

func time7() {
	// 设置第一个时间对象
	one := time.Now()
	// 设置第二个时间对象, 通过第一个时间对象加1小时获得
	two := one.Add(time.Hour)

	// 第二个时间对象调用Sub方法，减第一个时间对象，查看相差的时间
	res := two.Sub(one)
	fmt.Println(res)
}

func time8() {
	one := time.Now()
	two := one

	fmt.Println(one.Equal(two))
}

func time9() {
	one := time.Now()
	two := time.Now()

	fmt.Println(one.Before(two))
}

func time10() {
	one := time.Now()
	two := time.Now()

	fmt.Println(one.After(two))
}

func main() {
	// time1()

	// time2()

	// time3()

	// time4()

	// time5()

	// time6()

	// time7()

	time8()

	time9()

	time10()
}
