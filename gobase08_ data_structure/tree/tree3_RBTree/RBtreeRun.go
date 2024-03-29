package main

import "fmt"

//定义常量红黑
const (
	RED   = true
	BLACK = false
)

//数据接口
type Item interface {
	Less(than Item) bool
}

//红黑树结构
type RBNode struct {
	Left   *RBNode //左节点
	Right  *RBNode //右边节点
	Parent *RBNode //父亲节点
	Color  bool    //颜色
	//DataItem interface{} //数据
	Item //数据接口
}

//红黑树
type RBtree struct {
	NIL   *RBNode
	Root  *RBNode
	count uint
}

//比大小
type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type UInt32 uint32

func (x UInt32) Less(than Item) bool {
	return x < than.(UInt32)
}

type String string

func (x String) Less(than Item) bool {
	return x < than.(String)
}

//比大小
func less(x, y Item) bool {
	return x.Less(y)
}

//初始化内存
func NewRBTree() *RBtree {
	return new(RBtree).Init()
}

//初始化红黑树
func (rbt *RBtree) Init() *RBtree {
	node := &RBNode{nil, nil, nil, BLACK, nil}
	return &RBtree{node, node, 0}
}

//获取红黑树长度
func (rbt *RBtree) Len() uint {
	return rbt.count
}

//取得红黑树的极大值节点
func (rbt *RBtree) max(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for x.Right != rbt.NIL {
		x = x.Right
	}
	return x
}

//取得红黑树的极小值
func (rbt *RBtree) min(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for x.Left != rbt.NIL {
		x = x.Left
	}
	return x
}

//搜索红黑树
func (rbt *RBtree) search(x *RBNode) *RBNode {
	pnode := rbt.Root //根节点
	for pnode != rbt.NIL {
		if less(pnode.Item, x.Item) {
			pnode = pnode.Right
		} else if less(x.Item, pnode.Item) {
			pnode = pnode.Left
		} else {
			break //找到
		}
	}
	return pnode
}

//左旋转,x是要左旋的节点
func (rbt *RBtree) leftRotate(x *RBNode) {
	if x.Right == rbt.NIL {
		return //左旋转，逆时针，右孩子不可以为0
	}
	y := x.Right
	x.Right = y.Left //实现旋转的左旋
	if y.Left != rbt.NIL {
		y.Left.Parent = x //设定父亲节点
	}
	y.Parent = x.Parent //传递父节点
	if x.Parent == rbt.NIL {
		//根节点
		rbt.Root = y
	} else if x == x.Parent.Left { //x在根节点左边
		x.Parent.Left = y
	} else { //x在根节点右边
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

//右边旋转，左子树不可以为空
func (rbt *RBtree) rightRotate(x *RBNode) {
	if x.Left == nil {
		return
	}
	y := x.Left
	x.Left = y.Right
	if y.Right != rbt.NIL {
		y.Right.Parent = x //设置祖先
	}
	y.Parent = x.Parent //y保存x的父亲节点
	if x.Parent == rbt.NIL {
		rbt.Root = y
	} else if x == x.Parent.Left { //x小于根节点
		x.Parent.Left = y //父亲节点的孩子是x,改。父亲节点孩子y
	} else { //x大于根节点
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

//插入一条数据
func (rbt *RBtree) Insert(item Item) *RBNode {
	if item == nil {
		return nil
	}
	return rbt.insert(&RBNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item})
}

//插入
func (rbt *RBtree) insert(z *RBNode) *RBNode {

	//寻找插入位置
	x := rbt.Root
	y := rbt.NIL

	for x != rbt.NIL {
		y = x                     //备份位置，数据插入x,y之间
		if less(z.Item, x.Item) { //小于
			x = x.Left
		} else if less(x.Item, z.Item) { //大于
			x = x.Right
		} else { //相等
			return x //数据已经存在，无法插入
		}
	}
	z.Parent = y

	if y == rbt.NIL {
		rbt.Root = z
	} else if less(z.Item, y.Item) {
		y.Left = z //小于左边插入
	} else {
		y.Right = z //大于右边插入
	}
	rbt.count++
	rbt.insertFixup(z) //调整平衡
	return z
}

//插入之后，调整平衡
func (rbt *RBtree) insertFixup(z *RBNode) {
	for z.Parent.Color == RED { //一直循环下去，直到根节点
		if z.Parent == z.Parent.Parent.Left { //父亲节点在爷爷左边
			y := z.Parent.Parent.Right
			if y.Color == RED { //判断大伯节点红色，黑色

				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent //循环前进

			} else {
				if z == z.Parent.Right { //z比父亲小
					z = z.Parent
					rbt.leftRotate(z) //左旋

				} else { //z比父亲大
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.rightRotate(z.Parent.Parent)
				}
			}

		} else { // //父亲节点在爷爷右边
			y := z.Parent.Parent.Left //叔叔节点
			if y.Color == RED {       //判断大伯节点红色，黑色
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent //循环前进
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rbt.rightRotate(z)
				} else {
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.leftRotate(z.Parent.Parent)
				}
			}
		}
	}
	rbt.Root.Color = BLACK
}

//最小，最大，查找，修改，近似查找
func (rbt *RBtree) GetDepth() int {
	var getDeepth func(node *RBNode) int //函数

	//函数包含
	getDeepth = func(node *RBNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		var leftdeep int = getDeepth(node.Left)
		var rightdeep int = getDeepth(node.Right)
		if leftdeep > rightdeep {
			return leftdeep + 1
		} else {
			return rightdeep + 1
		}
	}
	return getDeepth(rbt.Root)
}

//近似查找
func (rbt *RBtree) searchle(x *RBNode) *RBNode {
	p := rbt.Root //根节点
	n := p        //备份根节点
	for n != rbt.NIL {
		if less(n.Item, x.Item) {
			p = n
			n = n.Right //大于
		} else if less(x.Item, n.Item) {
			p = n
			n = n.Left //小于
		} else {
			return n
			break //跳出循环
		}
	}
	if less(p.Item, x.Item) {
		return p
	}
	p = rbt.desuccessor(p) //近似处理
	return p
}

func (rbt *RBtree) successor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	if x.Right != rbt.NIL {
		return rbt.min(x.Right) //取得右边最小
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Right {
		x = y
		y = y.Parent
	}
	return y

}
func (rbt *RBtree) desuccessor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	if x.Left != rbt.NIL {
		return rbt.max(x.Left) //取得左边最大
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Left {
		x = y
		y = y.Parent
	}
	return y
}

//删除
func (rbt *RBtree) Delete(item Item) Item {
	if item == nil {
		return nil
	}
	return rbt.delete(&RBNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item}).Item
}
func (rbt *RBtree) delete(key *RBNode) *RBNode {
	z := rbt.search(key) //寻找要删除的节点
	if z == rbt.NIL {
		return rbt.NIL //无需删除
	}
	//新建节点下，x,y备份,夹逼
	var x *RBNode
	var y *RBNode

	//节点
	ret := &RBNode{rbt.NIL, rbt.NIL, rbt.NIL, z.Color, z.Item}

	if z.Left == rbt.NIL || z.Right == rbt.NIL {
		y = z //单节点，y,z重合
	} else {
		y = rbt.successor(z) //找到最接近的右边最小
	}

	if y.Left != rbt.NIL {
		x = y.Left
	} else {
		x = y.Right
	}
	x.Parent = y.Parent

	if y.Parent == rbt.NIL {
		rbt.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	if y != z {
		z.Item = y.Item
	}
	if y.Color == BLACK {
		rbt.deleteFixup(x)
	}
	rbt.count--
	return ret
}

//删除时的红黑修复需要考虑四种情况(下面的“X”指取代了被删节点位置的新节点)
// (1) X的兄弟节点是红色的：
//        |                              |
//       1●                             3●
//       / \             --------\      / \
// X-> 2●   ○3 <-brother --------/    1○   ●5
//         / \                        / \
//       4●   ●5                X-> 2●   ●4
//
// (2) X的兄弟节点是黑色的，而且兄弟节点的两个孩子都是黑色的：
//        |                              |
//       1○                         X-> 1○
//       / \             --------\      / \
// X-> 2●   ●3 <-brother --------/    2●   ○3
//         / \                            / \
//       4●   ●5                        4●   ●5
//
// (3) X的兄弟节点是黑色的，兄弟的左孩子是红色的，右孩子是黑色的：
//        |                              |
//       1○                             1○
//       / \             --------\      / \
// X-> 2●   ●3 <-brother --------/ X->2●   ●4
//         / \                              \
//       4○   ●5                             ○3
//                                            \
//                                             ●5
//
// (4) X的兄弟节点是黑色的，兄弟的右孩子是红色的：
//        |                              |
//       1○                             3○   X->root and loop while end
//       / \             --------\      / \
// X-> 2●   ●3 <-brother --------/    1●   ●5
//         / \                        / \
//       4○   ○5                    2●   ○4
//
//以上是兄弟节点在X右边时的情况，在X左边是取相反即可!
func (rbt *RBtree) deleteFixup(x *RBNode) {
	for x != rbt.Root && x.Color == BLACK {
		if x == x.Parent.Left { //x在左边
			w := x.Parent.Right //哥哥节点
			if w.Color == RED { //左边旋转
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent)
				w = x.Parent.Right //循环步骤
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent //循环条件

			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.rightRotate(w) //右旋转
					w = x.Parent.Right //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent)
				x = rbt.Root
			}
		} else { //x在右边
			w := x.Parent.Left  //左边节点
			if w.Color == RED { //左旋
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent)
				w = x.Parent.Right //循环步骤
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent //循环条件
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.leftRotate(w) //右旋转
					w = x.Parent.Left //循环条件
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.rightRotate(x.Parent)
				x = rbt.Root
			}
		}
	}
	x.Color = BLACK //循环到最后根节点，黑色
}

func main() {
	rbtree := NewRBTree()
	for i := 0; i < 10000000; i++ {
		rbtree.Insert(Int(i))
	}
	/*	for i := 0; i < 900000; i++ {
		//rbtree
		rbtree.Delete(Int(i))
	}*/
	fmt.Println(rbtree.GetDepth())
}
