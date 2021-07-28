package demo

import (
	"testing"
)

//1.功能测试
func TestGetArea(t *testing.T) { //函数方法相当于是测试用例
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

//2.压力测试
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
