package conf

import (
	"Go_Learn/Day_13/logAgent/model"
	"io/ioutil"
	"reflect"
	"strings"
)

// InitConf 初始化配置信息
func InitConf(file string, config *model.Config) (err error) {
	res, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	conf := strings.Split(string(res), "\n")

	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		field_t := v.Field(i).Type()
		field_v := v.Field(i)

		for j := 0; j < field_t.NumField(); j++ {
			tag := field_t.Field(j).Tag.Get("ini")

			for _, value := range conf {
				res := strings.Split(value, "=")
				if len(res) == 2 {
					key := strings.TrimSpace(res[0])
					val := strings.TrimSpace(res[1])
					if key == tag {
						switch field_t.Field(j).Type.Kind() {
						case reflect.String:
							field_v.Field(j).SetString(val)
						}
					}
				}
			}
		}

	}

	return
}
