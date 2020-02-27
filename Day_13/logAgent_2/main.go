package main

import (
	"Go_Learn/Day_13/logAgent_2/controller/conf"
	"Go_Learn/Day_13/logAgent_2/controller/etcd"
	"Go_Learn/Day_13/logAgent_2/controller/kafka"
	"Go_Learn/Day_13/logAgent_2/controller/log"
	"Go_Learn/Day_13/logAgent_2/model"
	"context"
	"encoding/json"
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

	// 初始化ETCD信息列表
	msg_list := map[string]*model.ETCDMsgMgr{}

	// 4. 从ETCD中获取数据
	res, err := etcd.GetMsg("xxx")
	if err != nil {
		fmt.Println("获取信息失败, err:", err)
		return
	}

	// 判断返回的信息中有没有值
	if res != nil {
		new(res, msg_list)
	}

	// 4.1 并监控数据变化
	rch := etcd.Cli.Watch(context.Background(), "xxx")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if fmt.Sprintf("%v", ev.Type) == "PUT" {
				// 序列化
				var msg model.ETCD_msg
				err = json.Unmarshal(ev.Kv.Value, &msg)
				if err != nil {
					fmt.Println("反序列化失败")
				}

				// 判断列表中有没有这个topic
				if res := msg_list[msg.Topic]; res != nil {
					if res.ETCD_msg.LogFile != msg.LogFile {
						// 关闭对应的goroutine
						res.Cancel()

						new(&msg, msg_list)

						continue
					}
				}

				new(&msg, msg_list)
			} else if fmt.Sprintf("%v", ev.Type) == "DELETE" {
				fmt.Printf("DELETE Key:%v\n", string(ev.Kv.Key))

			}
		}
	}
}

func new(msg *model.ETCD_msg, msg_list map[string]*model.ETCDMsgMgr) {
	ctx, cancel := context.WithCancel(context.Background())
	go run(ctx, msg)
	msg_list[msg.Topic] = &model.ETCDMsgMgr{msg, cancel}
}

func run(ctx context.Context, conf *model.ETCD_msg) {
	tails, err := log.ReadLog(conf.LogFile)
	if err != nil {
		fmt.Println("获取日志信息失败, err:", err)
	}

	for {
		select {
		case <-ctx.Done():
			// 执行关闭操作
			tails.Stop()
			tails.Cleanup()

			return
		case msg := <-tails.Lines:
			// go func() {

			// }()
			// 5. 将数据发送到kafka中
			pid, offset, err := kafka.SendToKafka(conf.Topic, msg.Text)
			if err != nil {
				fmt.Println("消息发送到kafka失败, err:", err)
				continue
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Tick(time.Second)
		}
	}
}
