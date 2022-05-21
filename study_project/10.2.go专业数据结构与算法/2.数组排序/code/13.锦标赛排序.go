package main

import (
	"fmt"
	"math"
)

type node struct {
	value int  //叶子的数据
	isok  bool //叶子状态是不是无穷大
	rank  int  //叶子的排序
}

func compareAndUp(tree *[]node, leftNode int) {
	rightNode := leftNode + 1

	// 除非左节点无效或者右节点有效并且比左节点大，否则就无脑选左节点
	if !(*tree)[leftNode].isok || ((*tree)[rightNode].isok && (*tree)[leftNode].value > (*tree)[rightNode].value) {
		mid := (leftNode - 1) / 2
		(*tree)[mid] = (*tree)[rightNode]
	} else {
		mid := (leftNode - 1) / 2
		(*tree)[mid] = (*tree)[leftNode]
	}

	//fmt.Println(tree)
}

//x^y //类型的处理
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func TreeSelectSort(arr []int) []int {

	var level int                         //树的层数
	var result = make([]int, 0, len(arr)) //保存最终结果

	for pow(2, level) < len(arr) {
		level++ //求出可以覆盖所有元素的层数
	}

	var leaf = pow(2, level)          //叶子数量
	var tree = make([]node, leaf*2-1) //树的节点数量
	//填充叶子节点
	for i := 0; i < len(arr); i++ {
		tree[leaf+i-1] = node{arr[i], true, i}
	}
	//进行对比
	for i := 0; i < level; i++ {
		nodeCount := pow(2, level-i) //每次处理降低一个层级/2
		for j := 0; j < nodeCount/2; j++ {
			leftnode := nodeCount - 1 + j*2
			compareAndUp(&tree, leftnode)
		}
	}
	fmt.Println("tree[0].value", tree[0].value)
	result = append(result, tree[0].value) //保存最顶端的最小数
	fmt.Println("res", result)
	//选出第一个以后，还有n-1个循环
	for t := 0; t < len(arr)-1; t++ {
		winnode := tree[0].rank + leaf - 1 //记录赢得的节点
		tree[winnode].isok = false         //修改成无穷大
		for i := 0; i < level; i++ {
			leftNode := winnode
			if winnode%2 == 0 { //处理奇数偶数
				leftNode = winnode - 1
			}
			compareAndUp(&tree, leftNode)

			winnode = (leftNode - 1) / 2 //保存中间节点
		}
		fmt.Println("tree[0].value", tree[0].value)
		result = append(result, tree[0].value) //保存最顶端的最小数
		fmt.Println("res", result)

	}

	return result

}

func main() {
	var length = 10
	var mymap = make(map[int]int, length)
	var obj []int
	//构造map
	for i := 0; i < length; i++ {
		mymap[i] = i //map哈希随机存储
	}
	for k, _ := range mymap {
		obj = append(obj, k) //叠加
	}

	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println("原始数组", arr)
	fmt.Println("sort", TreeSelectSort(arr))

}
