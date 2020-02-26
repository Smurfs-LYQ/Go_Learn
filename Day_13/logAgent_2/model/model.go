package model

import "context"

// Config 配置信息汇总
type Config struct {
	Kafka
	ETCD
}

// Kafka 配置信息
type Kafka struct {
	Address string `ini:"kafka_ip"`
}

// ETCD 配置信息
type ETCD struct {
	Address string `ini:"etcd_ip"`
}

// ETCD_msg ETCD返回的信息数据
type ETCD_msg struct {
	LogFile string `json:"logfile"`
	Topic   string `json:"topic"`
}

// ETCDMsgMgr ETCD返回的信息数据列表
type ETCDMsgMgr struct {
	*ETCD_msg
	Cancel context.CancelFunc
}
