package main

import (
	"Go_Learn/Day_06/project/03_logs/mylog"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// LogFileIni 设置日志文件信息
type LogFileIni struct {
	FileName string `conf:"file_name"`
	FilePath string `conf:"file_path"`
	MaxSize  int64  `conf:"max_size"`
}

// NewLogFileIni LogFileIni初始化函数
func NewLogFileIni(iniPath string, result interface{}) {
	// 拿到反射对象
	t := reflect.TypeOf(result)
	tElem := t.Elem()
	v := reflect.ValueOf(result).Elem()

	// 判断传入的是不是结构体指针
	if t.Kind() != reflect.Ptr && tElem.Kind() != reflect.Struct {
		panic("传入的必须是结构体指针类型")
	}

	// 读取文件
	file, err := ioutil.ReadFile(iniPath)
	if err != nil {
		panic(fmt.Errorf("文件打开失败"))
	}

	res := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, value := range res {
		result := strings.Split(value, "=")
		if len(result) != 2 {
			continue
		}

		key := strings.TrimSpace(result[0])
		val := strings.TrimSpace(result[1])

		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if tag == key {
				switch field.Type.Kind() {
				case reflect.String:
					v.Field(i).SetString(val)
				case reflect.Int64:
					result, _ := strconv.ParseInt(val, 10, 64)
					v.Field(i).SetInt(result)
				}
			}
		}
	}
}

// 定义接口的全局变量
var logger mylog.Mylogger

func main() {
	iniPath := "../ini/config.ini"
	var logConfig LogFileIni
	NewLogFileIni(iniPath, &logConfig)

	logger = mylog.NewFilelog("debug", logConfig.FileName, logConfig.FilePath, logConfig.MaxSize)
	for {
		logger.Debug("%s", "Debug日志")
		logger.Info("%s", "Info日志")
		logger.Warning("%s", "Warning日志")
		logger.Error("%s", "Error日志")
		logger.Fatel("%s", "Fatel日志")
		// time.Sleep(time.Second)
	}
}
