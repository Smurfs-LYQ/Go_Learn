package etcd

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	Cli *clientv3.Client
)

func InitETCD(ip []string, time time.Duration) (err error) {
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   ip,   // 添加节点
		DialTimeout: time, // 设置超时时间
	})

	return
}

// func Send(msg string) (err error) {

// }
