package Factory

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
