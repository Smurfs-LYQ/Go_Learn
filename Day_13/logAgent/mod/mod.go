package mod

// Config 配置信息
type Config struct {
	ETCD
	LogFile
}

// ETCD ETCD服务配置信息
type ETCD struct {
	IP string `ini:"kafka_ip"`
}

// LogFile 日志文件配置信息
type LogFile struct {
	Path string `ini:"log_path"`
}

type ETCD_msg struct {
	Path  string `json:"logfile"`
	Topic string `json:"topic"`
}
