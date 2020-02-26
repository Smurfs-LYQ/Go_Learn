package model

type Config struct {
	Kafka
	ETCD
}

type Kafka struct {
	Address string `ini:"kafka_ip"`
}

type ETCD struct {
	Address string `ini:"etcd_ip"`
}

type ETCD_msg struct {
	LogFile string `json:"logfile"`
	Topic   string `json:"topic"`
}
