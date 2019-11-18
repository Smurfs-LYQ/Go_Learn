package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
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
	t := reflect.TypeOf(result).Elem()
	v := reflect.ValueOf(result)

	for i := 0; i < t.NumField(); i++ {
		// res, ok := a[string]
		// 获取结构体字段的指定tag标签
		if res, ok := a[string(t.Field(i).Tag.Get("conf"))]; ok {
			fmt.Printf("%T %v\n", res, res)
			// 获取结构体字段的类型
			switch t.Field(i).Type.Kind() {
			case reflect.String:
				v.Elem().Field(i).SetString(res)
			case reflect.Int64:
				// result, _ := strconv.ParseInt(res, 10, 64)
				// v.Elem().Field(i).SetInt(result)
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
