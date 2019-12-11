package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("1,2,3", ",")      // 实际得到
	want := []string{"1", "2", "3"} // 希望得到的

	if !reflect.DeepEqual(got, want) {
		t.Errorf("期望得到的值 : %v，实际得到的值 : %v\n", got, want)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("1,2,3", ",")
	}
}

// 整个测试之前做的事和之后做的事
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前做的一些设置, 比如连接数据库
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retcode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做的一些拆卸工作
	os.Exit(retcode)                           // 退出测试
}
