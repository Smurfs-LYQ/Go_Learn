package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

// NSQ 生产者实例

var producer *nsq.Producer

// 初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()                    // 初始化一个空的配置信息
	producer, err = nsq.NewProducer(str, config) // 创建生产者 参数 1. 生产者名称 2. 生产者配置信息
	if err != nil {
		fmt.Println("创建生产者失败, err:", err)
		return err
	}
	return
}

func main() {
	nsqAddress := "127.0.0.1:4150"  // nsqd 的地址
	err := initProducer(nsqAddress) // 初始化生产者
	if err != nil {
		fmt.Println("初始化生产者失败, err:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) // 从标准输入读取
	// 循环读取标准输入
	for {
		data, err := reader.ReadString('\n') // 按照换行符分割获取用户输入
		if err != nil {
			fmt.Println("获取标准输入失败, err:", err)
			continue
		}
		data = strings.TrimSpace(data)    // 去除输入内容左右两侧的空格
		if strings.ToUpper(data) == "Q" { // 输入Q退出
			break
		}
		// 向指定的topic 'topic_demo' 发送数据
		/*
			如果没有这个topic，系统会自动创建
			参数:
				1. topic名称
				2. 信息数据
		*/
		err = producer.Publish("Test-1", []byte(data))
		if err != nil {
			fmt.Println("向消息队列发送消息失败, err:", err)
			continue
		}
	}
}
