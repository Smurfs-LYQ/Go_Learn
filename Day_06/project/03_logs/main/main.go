package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// LogFileIni 设置日志文件信息
type LogFileIni struct {
	fileName string `conf:"file_name"`
	filePath string `conf:"file_path"`
	maxSize  int64  `conf:"max_size"`
}

// NewLogFileIni LogFileIni初始化函数
func NewLogFileIni(iniPath string, result interface{}) {
	// 拿到反射对象
	t := reflect.TypeOf(result)
	tElem := t.Elem()
	// v := reflect.Value(result).Elem()

	// 判断传入的是不是结构体指针
	if t.Kind() != reflect.Ptr && tElem.Kind() != reflect.Struct {
		panic("传入的必须是结构体指针类型")
	}

	// 读取文件
	file, err := ioutil.ReadFile(iniPath)
	if err != nil {
		panic(fmt.Errorf("文件打开失败, 失败原因: %v\n", err))
	}

	res := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, v := range res {
		result := strings.Split(v, "=")
		if len(result) != 2 {
			continue
		}

		key := strings.TrimSpace(result[0])
		val := strings.TrimSpace(result[1])
		fmt.Println(key, val)
	}
}

func main() {
	iniPath := "../ini/config.ini"
	var logConfig LogFileIni
	NewLogFileIni(iniPath, &logConfig)
}
