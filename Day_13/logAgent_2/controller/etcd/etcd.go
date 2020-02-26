package etcd

import (
	"Go_Learn/Day_13/logAgent_2/model"
	"encoding/json"
	"time"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	Cli *clientv3.Client
)

// InitETCD 初始化ETCD连接
func InitETCD(address []string, time time.Duration) (err error) {
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   address, // 节点IP
		DialTimeout: time,    // 设置超时时间
	})

	return
}

// GetMsg 获取etcd中的信息
func GetMsg(key string) (msg model.ETCD_msg, err error) {
	ctx, concel := context.WithTimeout(context.Background(), time.Second*3)
	res, err := Cli.Get(ctx, key)
	concel()
	if err != nil {
		return
	}

	for _, v := range res.Kvs {
		err = json.Unmarshal(v.Value, &msg)
		if err != nil {
			return
		}
	}

	return
}
