package split

import (
	"reflect"
	"testing"
)

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如果需要在此执行: 测试之前的Setup")
	return func(t *testing.T) {
		t.Log("如果需要在此执行: 测试之后的Teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如果需要在此执行: 子测试之前的Setup")
	return func(t *testing.T) {
		t.Log("如果需要在此执行: 子测试之后的Teardown")
	}
}

// 使用方式
func TestSplit(t *testing.T) {
	// 定义test结构体
	type test struct {
		str  string
		sep  string
		want []string
	}

	// 测试用例使用map存储
	tests := map[string]test{
		"one": {str: "1,2,3", sep: ",", want: []string{"1", "2", "3"}},
		"two": {str: "1,2,3", sep: ",", want: []string{"1,2,3"}},
	}

	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行teardown操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run() 执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行teardown操作
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want: %#v, got:%#v\n", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("1,2,3", ",")
	}
}
