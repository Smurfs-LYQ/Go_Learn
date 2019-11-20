package main

import (
	"Go_Learn/Day_05/project_logs/mylog"
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
	// 拿到反射对象
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result).Elem()
	// 判断传入的对象是不是指针
	if t.Kind() != reflect.Ptr {
		panic("传入的对象不是结构体的指针")
	}

	// 因为传入的对象是指针，所以需要通过reflect包的Elem()方法拿到对象的值
	tElem := t.Elem()
	// 判断传入的对象底层类型是否为struct
	if tElem.Kind() != reflect.Struct {
		panic("传入的对象不是结构体")
	}

	// 打开文件
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("文件打开失败")
	}

	// 将读取到的内容转换成字符串类型，并将其切割
	file_res := strings.Split(string(file), "\n")

	for _, value := range file_res {
		v_slice := strings.Split(value, "=")
		// 判断如果长度不等于2则退出当前循环
		if len(v_slice) != 2 {
			continue
		}
		key := strings.TrimSpace(v_slice[0])
		val := strings.TrimSpace(v_slice[1])

		// 通过反射获取结构所有的字段
		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			// 拿到字段的tag标签
			tag := field.Tag.Get("conf")
			// 判断字段的tag和当前key(配置文件中等号左边的内容)是否相等
			if tag == key {
				// 判断字段的类型, 并给结构体反射对象设置值
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

var logger mylog.Mylogger

func main() {
	filepath := "../init/logs.conf"
	var c Config
	NewConfig(filepath, &c)

	logger = mylog.NewFilelog("debug", c.Filename, c.Filepath, c.Maxsize)
	defer logger.Die()

	for {
		logger.Debug("%s", "Debug日志")
		logger.Info("%s", "Info日志")
		logger.Warning("%s", "Warning日志")
		logger.Error("%s", "Error日志")
		logger.Fatel("%s", "Fatel日志")
	}
}
