package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// Config 解析日志库的配置文件
type Config struct {
	Filepath string `conf:"file_path"`
	Filename string `conf:"file_name"`
	Maxsize  int64  `conf:"max_size"`
}

// NewConfig Config结构体的构造函数
func NewConfig(filename string, result interface{}) {
	// 前提条件，result必须是一个ptr
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	if t.Kind() != reflect.Ptr {
		panic("必须传入Config对象地址\n")
	}

	tElem := t.Elem()
	if tElem.Kind() != reflect.Struct {
		panic("传入的Config对象必须是struct")
	}
	// 打开文件
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("文件打开失败, 失败原因: %v", err)
	}
	// 定义一个存conf的map
	// conf_list := make(map[string]interface{})
	// 将配置文件通过换行切割处理
	conf := strings.Split(string(file), "\n")
	for i := 0; i < len(conf); i++ {
		// 将通过上面测试的conf行进行切割
		a := strings.Split(conf[i], "=")
		// 如果长度小于2则代表不符合规范
		if len(a) != 2 {
			continue
		}
		key := strings.TrimSpace(a[0])
		val := strings.TrimSpace(a[1])
		// 利用反射给结构体赋值
		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			if field.Tag.Get("conf") == key {
				// 拿到每个字段的类型
				switch field.Type.Kind() {
				case reflect.String:
					v.Elem().Field(i).SetString(val)
				case reflect.Int64:
					value64, _ := strconv.ParseInt(val, 10, 64)
					v.Elem().Field(i).SetInt(value64)
				}
			}
		}
	}
}

func main() {
	filepath := "../init/logs.conf"
	var c Config
	NewConfig(filepath, &c)
	fmt.Println(c)
}
