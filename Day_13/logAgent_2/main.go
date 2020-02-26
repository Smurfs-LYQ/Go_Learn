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

type ETCDMsgMgr struct {
	*model.ETCD_msg
	Cancel context.CancelFunc
}

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

	list := []ETCDMsgMgr{}

	ctx, cancel := context.WithCancel(context.Background())
	go run(ctx, &res)
	list = append(list, ETCDMsgMgr{&res, cancel})

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

				var status bool

				// 循环已经运行的发送信息列表
				for k, v := range list {
					// 判断收到信息的topic与列表中的topic是否一致
					if msg.Topic == v.Topic {
						// 如果信息发生变化则关闭之前的goroutine 重新开启一个
						if v.LogFile != msg.LogFile {
							v.Cancel()
							// 从list中删除这个ETCDMsgMgr
							list = append(list[:k], list[k+1:]...)

							ctx, cancel := context.WithCancel(context.Background())
							go run(ctx, &msg)
							list = append(list, ETCDMsgMgr{&msg, cancel})
						}
						status = true
						break
					}
				}

				// 如果列表中没有对应的信息，那就新启动一个goroutine
				if !status {
					ctx, cancel := context.WithCancel(context.Background())
					go run(ctx, &msg)
					list = append(list, ETCDMsgMgr{&msg, cancel})
				}
			} else if fmt.Sprintf("%v", ev.Type) == "DELETE" {
				fmt.Printf("DELETE Key:%v\n", string(ev.Kv.Key))
				// for _, v := range list {
				// 	if v.Topic == key {
				// 		v.Cancel()
				// 		break
				// 	}
				// }
				// fmt.Println(list)
			}
		}
	}
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
			// 5. 将数据发送到kafka中
			pid, offset, err := kafka.SendToKafka(conf.Topic, msg.Text)
			if err != nil {
				fmt.Println("消息发送到kafka失败, err:", err)
				continue
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		}
	}
}
