package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// Config 解析日志库的配置文件
type Config struct {
	filepath string `conf:"file_path"`
	filename string `conf:"file_name"`
	maxsize  int64  `conf:"max_size"`
}

// NewConfig Config结构体的构造函数
func NewConfig(filename string, result interface{}) {
	/*
		打开文件
		一行一行读取内容，根据tag找结构体里面对应的字段
		找到了要赋值

	*/
	// 打开文件
	res, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("文件%s打开失败, 失败原因:%v", filename, err)
	}
	// 创建一个键值对map用来存放配置文件中的信息
	var a = make(map[string]interface{})
	for _, v := range strings.Split(string(res), "\n") {
		r := strings.Split(v, "=")
		a[r[0]] = r[1]
	}

	// 获取result变量的反射对象类型信息
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	for i := 0; i < t.NumField(); i++ {
		// 获取结构体字段的指定tag标签
		if res, ok := a[string(t.Field(i).Tag.Get("conf"))]; ok {
			// fmt.Println(v)
			// fmt.Printf("%T %v\n", t.Field(i).Type, t.Field(i).Type)
			// 获取结构体字段的类型并让他以string格式表示
			if t.Field(i).Type.String() == "string" {
				v.Field(i).SetString(reflect.ValueOf(res))
			} else {
				// fmt.Println("int")
				v.Field(i).SetInt(reflect.ValueOf(res))
			}
		}

	}
}

// SetConfig 给结构体初始化值
func SetConfig(filepath, filename string, maxsize int64) {

}

func main() {
	filepath := "../init/logs.conf"
	var c Config
	NewConfig(filepath, c)

}
