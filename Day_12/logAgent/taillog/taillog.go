package taillog

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

var (
	tails *tail.Tail
)

// Taillog_init taillog日志初始化
func Taillog_init() (err error) {
	filename := "logs/my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 如果日志文件被切割之后，会根据原文件名重新打开文件
		Follow:    true,                                 // 与上面的参数配合，设定是否根据 "原文件名"
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 设置打开文件之后从哪里开始读取数据，
		MustExist: false,                                // 设置如果日志文件不存在就报错
		Poll:      true,                                 // 使用轮循的方式
	}

	tails, err = tail.TailFile(filename, config)

	return
}

// Get_msg 获取信息
func Get_msg(log_ch chan<- string) {
	var (
		msg *tail.Tail
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines // chan tail.Line 从这里面接收读取到的数据
		if !ok {
			err = fmt.Errorf("获取信息失败 %v", time.Now().Format("2006-01-02 15:04:05.000"))
			// return
		}
		log_ch <- msg
	}
	// return
}
