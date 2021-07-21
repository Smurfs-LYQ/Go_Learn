package demo

import (
	"testing"
)

func TestOne(t *testing.T) {
	var tests = map[string]string{
		"1": "12321",
		"2": "油灯少灯油",
		"3": "LoL，,Lol",
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if Two(v) != true {
				t.Errorf("%s不是回文\n", v)
			}
		})
	}
}

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Two("12321")
	}
}
