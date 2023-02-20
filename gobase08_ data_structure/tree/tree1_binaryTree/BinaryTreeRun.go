package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type Node struct {
	Data  int   //数据
	Left  *Node //指向左边节点
	Right *Node //指向右边节点
}
type BinaryTree struct {
	Root *Node //根节点
	Size int   //数据的数量
}

//新建一个二叉树
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{}
	bst.Size = 0
	bst.Root = nil
	return bst
}

//获取二叉树大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

//判断是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

//根节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}
func (bst *BinaryTree) add(n *Node, data int) *Node { //插入节点 小的插入到左边，大的插入到右边
	if n == nil { //根节点
		bst.Size++
		//return &Node{Data:data}
		return &Node{data, nil, nil}
	}
	if data < n.Data {
		n.Left = bst.add(n.Left, data) //比我小，左边
	} else if data > n.Data {
		n.Right = bst.add(n.Right, data) //比我大，右边
	}
	return n
}

//判断数据是否存在
func (bst *BinaryTree) Isin(data int) bool {
	return bst.isin(bst.Root, data) //从根节点开始查找
}
func (bst *BinaryTree) isin(n *Node, data int) bool {
	if n == nil {
		return false //树是空树，找不到
	}
	if data == n.Data {
		return true
	} else if data < n.Data {
		return bst.isin(n.Left, data)
	} else {
		return bst.isin(n.Right, data)
	}
}

//找最大值
func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmax(bst.Root).Data //取得最大，为什么不用比较因为插入的时候就是最小右大
}
func (bst *BinaryTree) findmax(n *Node) *Node {
	if n.Right == nil {
		return n
	} else {
		return bst.findmax(n.Right)
	}
}

//找最小值
func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmin(bst.Root).Data //取得最大
}
func (bst *BinaryTree) findmin(n *Node) *Node {
	if n.Left == nil {
		return n
	} else {
		return bst.findmin(n.Left)
	}
}

//前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preorder(bst.Root)
}
func (bst *BinaryTree) preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	bst.preorder(node.Left)
	bst.preorder(node.Right)
}

//非递归前序遍历
func (bst *BinaryTree) PreOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			res = append(res, mybst.Data) //压入数据
			mystack.PushBack(mybst)       //首先左边压入栈中
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()     //挨个取出节点
			mybst = v.Value.(*Node) //实例化
			//res=append(res,mybst.Data)//压入数据
			mybst = mybst.Right //追加
			mystack.Remove(v)   //删除
		}
	}
	return res
}

//中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inorder(bst.Root)
}
func (bst *BinaryTree) inorder(node *Node) {
	if node == nil {
		return
	}

	bst.inorder(node.Left)
	fmt.Println(node.Data)
	bst.inorder(node.Right)
}

//非递归方式中序遍历
// 压入数据
//for  栈不为空
// if  else
func (bst *BinaryTree) InOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst) //首先左边压入栈中，root的左边可以一直到整个左边
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()           //挨个取出节点
			mybst = v.Value.(*Node)       //实例化
			res = append(res, mybst.Data) //压入数据
			mybst = mybst.Right           //追加
			mystack.Remove(v)             //删除
		}
	}
	return res
}

//后序遍历
func (bst *BinaryTree) PostOrder() {
	bst.postorder(bst.Root)
}
func (bst *BinaryTree) postorder(node *Node) {
	if node == nil {
		return
	}
	bst.postorder(node.Left)
	bst.postorder(node.Right)
	fmt.Println(node.Data)
}

//非递归后续遍历
func (bst *BinaryTree) PostOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	var PreVisited *Node  //提前访问的节点

	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst) //首先左边压入栈中
			mybst = mybst.Left      //左边
		}
		v := mystack.Back() //取出节点
		top := v.Value.(*Node)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && PreVisited == top.Left) || (PreVisited == top.Right) {
			res = append(res, top.Data) //压入数据
			PreVisited = top            //记录上一个节点
			mystack.Remove(v)           //处理完了在栈中删除

		} else {
			mybst = top.Right //右边循环
		}
	}
	return res
}

//树状显示二叉树
func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer                     //保存字符串
	bst.GenerateBSTstring(bst.Root, 0, &buffer) //调用函数实现遍历
	return buffer.String()
}
func (bst *BinaryTree) GenerateBSTstring(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		//buffer.WriteString(bst.GenerateDepthstring(depth)+"nil\n")//空节点
		return
	}
	//写入字符串，保存树的深度
	bst.GenerateBSTstring(node.Left, depth+1, buffer)
	buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Data) + "\n")
	bst.GenerateBSTstring(node.Right, depth+1, buffer)
}
func (bst *BinaryTree) GenerateDepthstring(depth int) string {
	var buffer bytes.Buffer //保存字符串
	for i := 0; i < depth; i++ {
		buffer.WriteString("--") //深度为0，1-- 2----
	}
	return buffer.String()
}

//删除最小
func (bst *BinaryTree) RemoveMin() int {
	ret := bst.FindMin()
	bst.Root = bst.removemin(bst.Root)
	return ret
}
func (bst *BinaryTree) removemin(n *Node) *Node {
	if n.Left == nil {
		//删除
		rightNode := n.Right //备份右边节点
		bst.Size--           //删除
		return rightNode
	}
	n.Left = bst.removemin(n.Left) //将rightNode赋值过去左边，左边就删掉了
	return n
}

//删除最大
func (bst *BinaryTree) RemoveMax() int {
	ret := bst.FindMax()
	bst.Root = bst.removemax(bst.Root)
	return ret
}
func (bst *BinaryTree) removemax(n *Node) *Node {
	if n.Right == nil {
		//删除
		leftNode := n.Left //备份右边节点
		bst.Size--         //删除
		return leftNode
	}
	n.Right = bst.removemax(n.Right) //将leftNode赋值过去左边，右边就删掉了
	return n
}

//删除数据
func (bst *BinaryTree) Remove(data int) {
	bst.Root = bst.remove(bst.Root, data)
}
func (bst *BinaryTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil //节点为空，不需要干活
	}
	if data < n.Data {
		n.Left = bst.remove(n.Left, data) //递归左边
		return n
	} else if data > n.Data {
		n.Right = bst.remove(n.Right, data) //递归右边
		return n
	} else {
		//处理左边为空
		if n.Left == nil {
			rightNode := n.Right //备份右边节点
			n.Right = nil
			bst.Size-- //删除
			return rightNode
		}
		//处理右边为空
		if n.Right == nil {
			leftNode := n.Left //备份右边节点
			n.Left = nil       //处理节点返回
			bst.Size--         //删除
			return leftNode
		}
		//左右节点都不为空
		oknode := bst.findmin(n.Right)        //找出比我小的节点顶替我
		oknode.Right = bst.removemin(n.Right) //6，7
		oknode.Left = n.Left                  //删除
		n.Left = nil                          //删除的清空
		n.Right = nil
		return oknode //实现删除
	}
}

//广度遍历，层级遍历
func (bst *BinaryTree) Levelshow() {
	bst.levelshow(bst.Root)
}
func (bst *BinaryTree) levelshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Front() //前面取出数据
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

//深度遍历
func (bst *BinaryTree) Stackshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Back() //前面取出数据 ,此时是栈
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

//最小公共祖先
func (bst *BinaryTree) FindlowerstAncestor(root *Node, nodea *Node, nodeb *Node) *Node {
	if root == nil {
		return nil
	}
	if root == nodea || root == nodeb {
		return root //有一个节点是根节点，
	}
	left := bst.FindlowerstAncestor(root.Left, nodea, nodeb)   //递归查找
	right := bst.FindlowerstAncestor(root.Right, nodea, nodeb) //递归查找
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

//求深度
func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	lengthleft := bst.GetDepth(root.Left)
	rightlength := bst.GetDepth(root.Right)
	if lengthleft > rightlength {
		return lengthleft + 1
	} else {
		return rightlength + 1
	}
}

//求节点总数
func (bst *BinaryTree) SumofTree(root *Node) int {
	if root == nil {
		return 0
	}
	leftnum := bst.SumofTree(root.Left)
	rightnum := bst.SumofTree(root.Right)
	return leftnum + rightnum + 1 //计算节点数量
}

func main() {
	bst := NewBinaryTree() //手动创建avl树
	node1 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node3 := &Node{6, nil, nil}
	node4 := &Node{1, nil, nil}
	node5 := &Node{3, nil, nil}
	node6 := &Node{5, nil, nil}
	node7 := &Node{7, nil, nil}
	bst.Root = node1
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	bst.Size = 7

	fmt.Println(bst.Isin(3)) //数据是否存在
	fmt.Println(bst.FindMax(), "max")
	fmt.Println(bst.FindMin(), "min")
	fmt.Println("中序----------------------")
	bst.InOrder()
	fmt.Println("前序----------------------")
	bst.PreOrder()
	fmt.Println("后序----------------------")
	bst.PostOrder()
	//非递归遍历
	fmt.Println("非递归中序----------------------")
	fmt.Println(bst.InOrderNoRecursion())
	fmt.Println("非递归前序----------------------")
	fmt.Println(bst.PreOrderNoRecursion())
	fmt.Println("非递归后序----------------------")
	fmt.Println(bst.PostOrderNoRecursion())
	fmt.Println("level----------------------")
	bst.Levelshow()         //广度遍历，层级遍历
	bst.Stackshow(bst.Root) //深度遍历

	fmt.Println("last ----------------------")
	fmt.Println(bst.String())

	nodelast := bst.FindlowerstAncestor(bst.Root, node6, node7) //最小公共祖先
	fmt.Println("最小公共祖先是：", nodelast)
	fmt.Println(bst.GetDepth(bst.Root))  //求深度
	fmt.Println(bst.SumofTree(bst.Root)) //求节点总数

	fmt.Println(bst.RemoveMin())
	fmt.Println(bst.RemoveMax())
	bst.Remove(4)
}
