package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filePath := "../ini/one.ini"

	res, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("文件读取失败, 错误信息: %v\n", err))
	}

	file := strings.Split(string(res), "\n")
	for k, v := range file {
		// 判断是不是注释
		if strings.HasPrefix(v, ";") || strings.HasPrefix(v, "#") {
			continue
		}

		// 判断是不是不同服务的配置文件
		if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
			fmt.Printf("\n%s 信息\n", strings.Trim(v, "[]"))
			continue
		}

		// 进行字符串切割
		vSplit := strings.Split(v, "=")
		if len(vSplit) != 2 {
			fmt.Printf("这一行配置信息不合法，行号: %d\n", k+1)
			continue
		}
		key := strings.TrimSpace(vSplit[0])
		val := strings.TrimSpace(vSplit[1])
		fmt.Printf("%s\t%s\n", key, val)
	}
}
