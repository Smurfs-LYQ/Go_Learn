package main

import (
	"Go_Learn/Day_13/logAgent/etcd"
	"Go_Learn/Day_13/logAgent/log"
	"Go_Learn/Day_13/logAgent/mod"
	"Go_Learn/Day_13/logAgent/taillog"
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	var config mod.Config

	// 1. 打开配置文件，获取配置信息
	err := log.InitLogs("conf/Config.ini", &config)
	if err != nil {
		fmt.Println("初始化配置信息失败, err:", err)
		return
	}
	fmt.Println("初始化配置信息成功")

	// 2. 初始化etcd
	err = etcd.InitETCD(strings.Split(config.ETCD.IP, ","))
	if err != nil {
		fmt.Println("连接ETCD失败, err:", err)
		return
	}
	fmt.Println("连接ETCD成功")
	defer etcd.Cli.Close()

	// 3. 获取日志信息
	err = taillog.InitTaillog(config.LogFile.Path)
	if err != nil {
		fmt.Println("获取日志信息服务启动失败, err:", err)
		return
	}
	fmt.Println("获取日志信息服务启动成功")

	// 4. 发送到etcd
	run()
}

func run() {
	var con int
	for msg := range taillog.Tails.Lines {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		_, err := etcd.Cli.Put(ctx, fmt.Sprintf("log_%d", con), fmt.Sprintf("%v", msg.Text))
		cancel()
		if err != nil {
			fmt.Println("信息发送到ETCD失败, err:", err)
		}
		fmt.Println("信息发送成功")
		con++
	}
}
