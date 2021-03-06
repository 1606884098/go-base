package BST

//二叉树结构
type  node struct{
	data int
	left *node
	right *node
}
type BST struct{
	root *node
	size int //树根
}

type BSTfunc interface {
	Getsize()
	Add()
	AddNode()
	Contains()
	BinSearch()
	PreOrder()
	InOrder()
	PostOrder()
	Min()
	Max()
	Remove()
	makeTree()
	GetDepth()
}

func NEWBST() *BST{ //创建树根
	bst:=&BST{}
	bst.root=nil
	bst.size=0
	return bst
}

