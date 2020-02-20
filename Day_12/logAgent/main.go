package main

import (
	"Go_Learn/Day_12/logAgent/kafka"
	"Go_Learn/Day_12/logAgent/taillog"
	"fmt"
)

func main() {
	// 初始化kafka连接
	err := kafka.Kafka_init()
	if err != nil {
		panic(fmt.Sprintf("kafka启动失败, err:%v\n", err))
	}
	// 打开日志文件准备收集日志
	err = taillog.Taillog_init()
	if err != nil {
		panic(fmt.Sprintf("日志收集系统启动失败, err:%v\n", err))
	}

	log_ch := make(chan string, 100)

	// 获取信息
	go taillog.Get_msg(log_ch<-)

	select {
	case msg := <- log_ch:
		fmt.Printf("%v %T\n", msg, msg)
		// 发送信息
		err := kafka.Send("test", msg)
		if err != nil {
			fmt.Printf("信息发送失败, err:%v\n", err)
		}
	default:
	}
}
