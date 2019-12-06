package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")      // 实际的结果
	want := []string{"a", "b", "c"} // 期待的结果

	// 深度判断: 先判断类型是否一致，再判断里面的元素是否都一致, 返回一个bool
	/*
		res := reflect.DeepEqual(got, want)
		fmt.Println(res)
	*/

	// 直接在if语句块中设置值，然后对此值进行判断，在这里声明的变量只在if语句块中生效
	if res := reflect.DeepEqual(got, want); !res {
		fmt.Println(t.Name(), "测试不通过", res) // t.Name() 用于获取当前正在测试的函数名
		t.Fail()                            // t.Fail() 用于宣告测试失败
	}
}

func TestNoneSplit(t *testing.T) {
	t.Log("如果字符串中不包含分隔符，测试结果是否正确") // 记录一些运行时的日志

	got := Split("a:b:c", "*")
	want := []string{"a", "b", "c"}

	if res := reflect.DeepEqual(got, want); !res {
		t.Fatalf("期望得到: %v, 实际得到: %v\n", want, got) // t.Fatalf() fatal报错
	}
}
