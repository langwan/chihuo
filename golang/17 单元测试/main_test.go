package main

import "testing"

func Test_Add(t *testing.T) {
	result := Add(1, 2)
	if result == 3 {
		t.Logf("ok")
	} else {
		t.Fail()
	}
}

func Test_AddAll(t *testing.T) {
	tests := []struct {
		a int
		b int
		r int
	}{{1, 2, 3}, {2, 2, 4}, {3, 3, 6}}

	for _, test := range tests {
		if output := Add(test.a, test.b); output != test.r {
			t.Error("失败 参数 ({}, {}) 结果 {}", test.a, test.b, test.r)
		}
	}
}
