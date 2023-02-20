package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

/* 打桩(stubs)
   当 Get() 的参数为 Tom，则返回 error，这称之为打桩(stub)，有明确的参数和返回值是最简
   单打桩方式。除此之外，检测调用次数、调用顺序，动态设置返回值等方式也经常使用。
   参数(Eq, Any, Not, Nil)
   • Eq(value) 表示与 value 等价的值。
   • Any() 可以用来表示任意的入参。
   • Not(value) 用来表示非 value 以外的值。
   • Nil() 表示 None 值
   返回值(Return, DoAndReturn)
   • Return 返回确定的值
   • Do Mock 方法被调用时，要执行的操作吗，忽略返回值。
   • DoAndReturn 可以动态地控制返回值。
   调用次数(Times)
   • Times() 断言 Mock 方法被调用的次数。
   • MaxTimes() 最大次数。
   • MinTimes() 最小次数。
   • AnyTimes() 任意次数（包括 0 次）。
   调用顺序(InOrder)
*/
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用
	//mockgen -source=db.go -destination=db_mock.go -package=main 自动生成的
	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
