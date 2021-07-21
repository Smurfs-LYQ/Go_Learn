package main

import (
	"Go_Learn/Day_12/logAgent/kafka"
	"Go_Learn/Day_12/logAgent/taillog"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type config struct {
	IP       string `conf:"ip"`
	Topic    string   `conf:"topic"`
	Log_file string   `conf:"log_file"`
}

func run() {
	log_ch := make(chan string, 100)

	// 获取信息
	go taillog.Get_msg(log_ch)

	/*
		for {
			select {
			case msg := <-log_ch:
				// 发送信息
				err := kafka.Send("test", msg)
				if err != nil {
					fmt.Printf("信息发送失败, err:%v\n", err)
				}
			default:
			}
		}
	*/

	for msg := range log_ch {
		err := kafka.Send("test", msg)
		if err != nil {
			fmt.Printf("信息发送失败，err:%v\n", err)
		}
	}
}

func new_config(config *config) {
	// 通过反射获取结构体的tag信息
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()

	// 查询配置信息
	conf_res, err := ioutil.ReadFile("config/config.conf")
	if err != nil {
		panic(fmt.Sprintf("配置文件加载失败"))
	}

	res := strings.Split(string(conf_res), "\n")
	for _, val := range res {
		res := strings.Split(strings.TrimSpace(val), "=")
		if len(res) >= 2 {
			for i := 0; i < t.NumField(); i++ {
				tag := t.Field(i).Tag.Get("conf")

				if strings.TrimSpace(res[0]) == tag {
					switch t.Field(i).Type.Kind() {
					case reflect.String:
						v.Field(i).SetString(strings.TrimSpace(res[1]))
					}
					break
				}
			}
		}
	}
}

func main() {
	// 初始化配置信息
	var config config
	new_config(&config)

	ip_list := strings.Split(config.IP, ",")

	// 初始化kafka连接
	err := kafka.Kafka_init(ip_list[:len(ip_list)-1])
	//err := kafka.Kafka_init([]string{"192.168.1.101:9092"})
	if err != nil {
		panic(fmt.Sprintf("kafka启动失败, err:%v\n", err))
	}
	fmt.Println("kafka连接成功")

	defer kafka.Client.Close()
	// 打开日志文件准备收集日志
	err = taillog.Taillog_init(config.Log_file)
	if err != nil {
		panic(fmt.Sprintf("日志收集系统启动失败, err:%v\n", err))
	}
	fmt.Println("日志收集系统启动成功")

	run()
}
