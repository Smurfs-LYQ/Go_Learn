package split

import (
	"testing"
)

func BenchmarkSplitParallel(b *testing.B) { // Parallel 用来这个函数为并行测试函数
	// b.SetParallelism(1) // 设置使用的CPU核心数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("1,2,3", ",")
		}
	})
}
