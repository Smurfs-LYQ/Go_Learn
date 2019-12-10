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

// 将多个测试用例放到一起
