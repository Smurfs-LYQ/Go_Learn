package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

// NSQ 消费者实例

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg: %v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

// 初始化消费者
/*
	参数:
		topic名称
		channel名称 消费者订阅了这个channel才会创建这个channel
		节点地址，可以是lookupd地址也可以是nsqd地址
*/
func InitConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()                         // 创建一个配置信息变量
	config.LookupdPollInterval = 15 * time.Second     // 设置nsqlookupd每15秒检查一次有没有新的nsqd节点加入进来
	c, err := nsq.NewConsumer(topic, channel, config) // 创建一个消费者
	if err != nil {
		fmt.Println("创建消费者失败, err:", err)
		return err
	}
	// 创建一个结构体对象
	consumer := &MyHandler{
		Title: "Smurfs的格格巫",
	}

	// 将结构体放入到 消费者(c) 的处理程序中
	c.AddHandler(consumer)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询获得要调用的那个NSQD
		return err
	}

	return
}

func main() {
	err := InitConsumer("Test-1", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Println("初始化消费者失败, err:", err)
		return
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}
