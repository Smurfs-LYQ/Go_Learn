package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// LogFileIni 设置日志文件信息
type LogFileIni struct {
	fileName string
	filePath string
	maxSize  int
}

func getIni() {
	// 打开文件
	file_res, err := ioutil.ReadFile("../ini/config.ini")
	if err != nil {
		panic(fmt.Sprintf("文件打开失败, %v\n", err))
	}
	res := strings.Split(string(file_res), "\n")
	for _, v := range res {
		value := strings.Split(v, "=")
		// 判断返回值长度
		if len(value) != 2 {
			continue
		}
		// 去除字符串左右的空格
		key := strings.TrimSpace(value[0])
		val := strings.TrimSpace(value[1])
		fmt.Println(key, val)
	}
}

func main() {
	getIni()
}
