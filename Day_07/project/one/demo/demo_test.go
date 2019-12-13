package demo

import (
	"testing"
)

func TestOne(t *testing.T) {
	str := "12321"

	var tests = map[string]string{
		"1": "12321",
		"2": "123521",
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if One(v) != true {
				t.Errorf("%s不是回文\n", str)
			}
		})
	}
}

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		One("12321")
	}
}
