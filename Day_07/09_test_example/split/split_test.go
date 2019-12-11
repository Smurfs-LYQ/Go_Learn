package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		str  string
		sep  string
		want []string
	}

	tests := map[string]test{
		"one": test{"1,2,3", ",", []string{"1", "2", "3"}},
		"two": test{"1,2,3", ",", []string{"1,2,3"}},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("你想得到的: %v, 实际得到的:%v\n", v.want, got)
			}
		})
	}
}

// 示例函数
func ExampleSplit() {
	fmt.Println(Split("1,2,3", ","))
	// Output:
	// [1 2 3]
}
