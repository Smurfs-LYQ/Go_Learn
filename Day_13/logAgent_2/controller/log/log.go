package log

import "github.com/hpcloud/tail"

func ReadLog(file string) (tails *tail.Tail, err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 如果日志文件被切割之后，会根据原文件名重新打开文件
		Follow:    true,                                 // 与上面的参数配合，设定是否根据 "原文件名"
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 设置打开文件之后从哪里开始读取数据，
		MustExist: false,                                // 设置如果日志文件不存在就报错
		Poll:      true,                                 // 使用轮循的方式
	}

	tails, err = tail.TailFile(file, config)
	if err != nil {
		return
	}

	return
}
