package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// 定义一个存放测试数据的结构体
	type test struct {
		str  string   // 字符串
		sep  string   // 切割字符
		want []string // 期望得到的值
	}

	// 创建一个存放多个测试用例的map
	var tests = map[string]test{
		"normal": test{"1,2,3", ",", []string{"1", "2", "3"}},
		"none":   test{"1:2:3", ":", []string{"1", "2", "3"}},
		"multi":  test{"1:2:3", ":2:", []string{"1", "2", "3"}},
	}

	// 循环调用测试用例
	for k, v := range tests {
		t.Run(k, func(t *testing.T) { // 使用t.Run()执行子测试
			res := Split(v.str, v.sep) // 将测试用例中的数据放入到测试的函数中
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("期望找到: %v, 实际得到: %v\n", v.want, res)
			}
		})
	}
}
