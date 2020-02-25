package log

import (
	"Go_Learn/Day_13/logAgent/mod"
	"io/ioutil"
	"reflect"
	"strings"
)

func InitLogs(file string, Config *mod.Config) (err error) {
	v := reflect.ValueOf(Config).Elem()

	// 打开日志文件
	res, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	// 将日志分行切割
	logVal := strings.Split(string(res), "\n")

	// 循环遍历config结构体
	for i := 0; i < v.NumField(); i++ {
		field_v := v.Field(i)

		// 循环遍历config结构体 中的子结构体
		for j := 0; j < field_v.Type().NumField(); j++ {
			// 获取子结构体field的tag
			tag := field_v.Type().Field(j).Tag.Get("ini")

			// 循环读取文件信息
			for _, val := range logVal {
				res := strings.Split(val, "=")
				if len(res) == 2 {
					// 判断tag是否与配置信息名一直
					if tag == strings.TrimSpace(res[0]) {
						// 遍历子结构体元素的类型
						switch field_v.Type().Field(j).Type.Kind() {
						case reflect.String:
							// 设置值
							field_v.Field(j).SetString(strings.TrimSpace(res[1]))
						}
						break
					}
				}
			}
		}
	}
	return
}
