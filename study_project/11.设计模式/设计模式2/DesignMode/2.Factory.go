package main

import "fmt"

// A X B
//X操作
//A B  操作数
// x=+  A+B  x=-  A-b  x=%  A%B
// left, right,

//实际运行类的接口
type Operator interface {
	Setleft(int)
	SetRight(int)
	Result() int
}

//工厂接口
type OperatorFactory interface {
	Create() Operator
}

//数据
type OperatorBase struct {
	left, right int
}

//赋值
func (op *OperatorBase) Setleft(left int) {
	op.left = left
}
func (op *OperatorBase) SetRight(right int) {
	op.right = right
}

//操作的抽象
type PlusOperatorFactory struct{}

//操作类中包含操作数
type PlusOperator struct {
	*OperatorBase
}

//实际运行
func (o PlusOperator) Result() int {
	return o.right + o.left
}
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}

//操作的抽象
type SubOperatorFactory struct{}

//操作类中包含操作数
type SubOperator struct {
	*OperatorBase
}

//实际运行
func (o SubOperator) Result() int {
	return o.left - o.right
}
func (SubOperatorFactory) Create() Operator {
	return &SubOperator{&OperatorBase{}}
}

func main21() {
	var fac OperatorFactory
	//fac=PlusOperatorFactory{}
	fac = SubOperatorFactory{}
	op := fac.Create()
	op.Setleft(20)
	op.SetRight(10)

	fmt.Println(op.Result())

}
