package mylogger

import "testing"

// 单元测试 测试方法: 到该目录下运行"go test -v"
func TestConstLevel(t *testing.T) {
	t.Logf("%v %T\n", DebugLevel, DebugLevel)
	t.Logf("%v %T\n", FatelLevel, FatelLevel)
}
