package main

import (
	"errors"
	"fmt"
)

// writer 定义一个接口
type writer interface {
	Write([]byte) error
}

type Student struct {
	name string
}

func (s Student) Write(one []byte) err error {
	if s.name != "" {
		fmt.Println(s.name, "正在写: ", string(one))
	} else {
		err = errors.New("这个学生没备注名字")
		return
	}
	return
}

func main() {
	var 
}
