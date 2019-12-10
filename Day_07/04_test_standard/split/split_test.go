package split

import "testing"

// 基准测试
func BenchmarkSplit(b *testing.B) {
	/*
		// 创建用于存储测试用例信息的结构体
		type test struct {
			str  string
			sep  string
			want []string
		}

		// 将测试用例信息存储map中
		var tests = map[string]test{
			"normal": {"1,2,3", ",", []string{"1", "2", "3"}},
			"none":   {"1:2:3", ":", []string{"1", "2", "3"}},
			"multi":  {"1,2,3", ",2,", []string{"1", "2", "3"}},
		}

		for k, v := range tests {
			t.Run(k, func(t *testing.T) {
				got := Split(v.str, v.sep)
				if !reflect.DeepEqual(got, v.want) {
					t.Errorf("期望找到: %v, 实际得到: %v\n", v.want, got)
				}
			})
		}
	*/
	// b.Log("这是一个基准测试")

	for i := 0; i < b.N; i++ {
		Split("1,2,3", ",")
	}
}
