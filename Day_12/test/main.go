package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

// producer 发送消息
func producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.1.101:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

// consumer 接收消息
func consumer() {
	consumer, err := sarama.NewConsumer([]string{"192.168.1.101:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("test") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有发的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		for msg := range pc.Messages() {
			fmt.Printf("Partition: %d Offset: %d Key: %v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}

		// 异步从每个分区消费信息
		// go func(sarama.PartitionConsumer) {
		// 	for msg := range pc.Messages() {
		// 		fmt.Printf("Partition: %d Offset: %d Key: %v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		// 	}
		// }(pc)
	}
}

func main() {
	// producer()

	consumer()
}
