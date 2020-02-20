package kafka

import (
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

// Kafka_init 初始化kafka连接
func Kafka_init() (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.1.101:9092"}, config)
	if err != nil {
		return
	}

	fmt.Println("kafka连接成功")
	return
}

// Send 发送信息
func Send(topic, val string) (err error) {

	// 构建一个信息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic // 设置信息主体
	msg.Value = sarama.StringEncoder(val)

	// 发送信息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return
	}

	fmt.Printf("信息发送成功, pid:%v offset:%v\n", pid, offset)

	return
}
