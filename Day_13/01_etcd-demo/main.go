package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.101:2379"}, // 添加节点
		DialTimeout: time.Second * 3,                // 设置连接超时时间
	})

	if err != nil {
		fmt.Printf("connect to etcd faield, err: %v\n", err)
		return
	}
	defer cli.Close()

	fmt.Println("connect to etcd success")

	// put
	ctx, cannel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "smurfs", "格格巫")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd faield, err: %v\n", err)
		return
	}

	// get
	ctx, cannel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "smurfs")
	cannel()
	if err != nil {
		fmt.Printf("get to etcd faield, err: %v\n", err)
		return
	}
	for k, v := range resp.Kvs {
		fmt.Println(k, v)
	}
}
