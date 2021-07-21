package split

import (
	"reflect"
	"testing"
)

// 单个测试用例
func TestSplit(t *testing.T) {
	// got := Split("1,2,3", ",")    // 实际得到的值
	got := Split("1,2,3", ",2,")    // 实际得到的值
	want := []string{"1", "2", "3"} // 期待得到的值
	t.Log("测试返回值与期待值是否一致")

	// 深度对比
	if res := reflect.DeepEqual(got, want); !res {
		t.Errorf("信息不一致\ngot = %v\nwant = %v\n", got, want)
	}
}

// 将多个测试用例放到一起组成 测试组
func TestSplitGroup(t *testing.T) {
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
		res := Split(v.str, v.sep) // 将测试用例中的数据放入到测试的函数中
		if !reflect.DeepEqual(res, v.want) {
			t.Errorf("测试用例: %v失败, 期望找到: %v, 实际得到: %v\n", k, v.want, res)
		}
	}
}
