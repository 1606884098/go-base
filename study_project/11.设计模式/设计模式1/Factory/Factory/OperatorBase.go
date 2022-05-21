package Factory

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
