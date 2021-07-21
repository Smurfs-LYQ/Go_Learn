package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 如果日志文件被切割之后，会根据原文件名重新打开文件
		Follow:    true,                                 // 与上面的参数配合，设定是否根据 "原文件名"
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 设置打开文件之后从哪里开始读取数据，
		MustExist: false,                                // 设置如果日志文件不存在就报错
		Poll:      true,                                 // 使用轮循的方式
	}
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tail %s faield, err: %v\n", filename, err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines // chan tail.Line 从这里面接收读取到的数据
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text) // 拿到信息的文本
	}
}
