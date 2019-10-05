package one

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Suiji_do struct {
	Now_id int
}

func (this *Suiji_do) Suiji(lists []*Objs) (obj *Objs, err error) {
	rand.Seed(time.Now().Unix())

	len := len(lists)

	if len <= 0 {
		err = errors.New("没有可用主机")
		return
	}

	num := rand.Intn(len - 1)
	obj = lists[num]
	return
}

func (this *Suiji_do) Xunhuan(lists []*Objs) (obj *Objs, err error) {
	// func (this *Xunhuan_do) Xunhuan(lists []*Objs) (obj *Objs, err error) {
	len := len(lists)

	if len <= 0 {
		err = errors.New("没有可用主机")
		return
	}

	obj = lists[this.Now_id]
	this.Now_id = (this.Now_id + 1) % len

	return
}

func Do() {
	// 创建主机列表
	lists := []*Objs{}
	for i := 0; i < 3; i++ {
		var one Objs = Objs{
			host: fmt.Sprintf("192.168.1.%d", i),
			port: i,
		}
		lists = append(lists, &one)
	}

	// 实例化接口
	var T1 One

	T1 = &Suiji_do{}
	for i := 0; i < 5; i++ {
		res, err := T1.Suiji(lists)
		if err != nil {
			fmt.Println("调用主机出错, 错误信息: ", err)
			continue
		} else {
			fmt.Println("随机调度：", res, "IP: ", res.host, "端口: ", res.port)
		}
		time.Sleep(time.Second)
	}

	fmt.Println("-------------------------------------------------------")

	for i := 0; i < 6; i++ {
		res, err := T1.Xunhuan(lists)
		if err != nil {
			fmt.Println("调用主机出错, 错误信息: ", err)
			continue
		} else {
			fmt.Println("随机调度：", res, "IP: ", res.host, "端口: ", res.port)
		}
	}
}
