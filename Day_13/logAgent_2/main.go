package main

import (
	"Go_Learn/Day_13/logAgent_2/controller/conf"
	"Go_Learn/Day_13/logAgent_2/controller/etcd"
	"Go_Learn/Day_13/logAgent_2/controller/kafka"
	"Go_Learn/Day_13/logAgent_2/model"
	"fmt"
	"strings"
	"time"
)

var config model.Config

func main() {
	// 1. 初始化配置文件
	err := conf.InitConf("conf/config.ini", &config)
	if err != nil {
		fmt.Println("初始化配置文件失败, err:", err)
		return
	}
	fmt.Println("配置文件初始化成功")

	// 2. 初始化ETCD连接
	err = etcd.InitETCD(strings.Split(config.ETCD.Address, ","), time.Second*3)
	if err != nil {
		fmt.Println("ETCD初始化失败, err:", err)
		return
	}
	fmt.Println("ETCD初始化成功")
	defer etcd.Cli.Close()

	// 3. 初始化kafka连接
	err = kafka.InitKafka(strings.Split(config.Kafka.Address, ","))
	if err != nil {
		fmt.Println("kafka初始化失败, err:", err)
		return
	}
	fmt.Println("kafka初始化成功")
	defer kafka.Client.Close()

	// 4. 从ETCD中获取数据
	res, err := etcd.GetMsg("xxx")
	if err != nil {
		fmt.Println("获取信息失败, err:", err)
		return
	}
	fmt.Println(res)

	// 4.1 并监控数据变化

	// 5. 将数据发送到kafka中
}
