package main

import "testing"

func TestSum(t *testing.T) {
	set := []int{17, 23, 100, 76, 55}
	expected := 271
	actual := Sum(set)
	if actual != expected {
		// 打印失败信息
		t.Errorf("Expect %d, but got %d!", expected, actual)
	}else {
		// 打印成功信息
		t.Logf("ok")
	}
}

func Benchmark_Sum(b *testing.B) {
	set := []int{17, 23, 100, 76, 55}
	// 压力测试循环体内要使用testing.B.N,以测试正常运行
	for i := 0; i < b.N; i++ {
		Sum(set)
	}
}