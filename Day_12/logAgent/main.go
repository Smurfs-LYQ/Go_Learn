package main

import (
	"Go_Learn/Day_12/logAgent/kafka"
	"Go_Learn/Day_12/logAgent/taillog"
	"fmt"
)

func run() {
	log_ch := make(chan string, 100)

	// 获取信息
	go taillog.Get_msg(log_ch)

	/*
	for {
		select {
		case msg := <-log_ch:
			// 发送信息
			err := kafka.Send("test", msg)
			if err != nil {
				fmt.Printf("信息发送失败, err:%v\n", err)
			}
		default:
		}
	}
	*/

	for msg := range log_ch {
		err := kafka.Send("test", msg)
		if err != nil {
			fmt.Printf("信息发送失败，err:%v\n", err)
		}
	}
}

func main() {
	// 初始化kafka连接
	err := kafka.Kafka_init([]string{"192.168.1.101:9092"})
	if err != nil {
		panic(fmt.Sprintf("kafka启动失败, err:%v\n", err))
	}
	fmt.Println("kafka连接成功")

	defer kafka.Client.Close()
	// 打开日志文件准备收集日志
	err = taillog.Taillog_init("./logs/my.log")
	if err != nil {
		panic(fmt.Sprintf("日志收集系统启动失败, err:%v\n", err))
	}
	fmt.Println("日志收集系统启动成功")

	run()
}
