package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	res, err := ioutil.ReadFile("../10_read_file/xx.txt")
	if err != nil {
		fmt.Println("文件读取失败，错误信息: ", err)
		return
	}
	fmt.Println(string(res))
}
