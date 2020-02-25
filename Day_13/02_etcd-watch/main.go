package main

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.1.101:2379"},
	})
	if err != nil {
		fmt.Printf("connect to etcd faield, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// watch key:smurfs change
	// 一直监视 smurfs 这个key的变化(新增、修改、删除)
	rch := cli.Watch(context.Background(), "smurfs") // <- chan WatchResponse
	// 从通道尝试取值(监视的信息)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%v Key:%v Value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
