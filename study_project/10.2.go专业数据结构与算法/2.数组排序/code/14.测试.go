package main

import (
	"fmt"
	"math"
)

// 声明一个节点的结构体，包含节点数值大小和是否需要参与比较
type node struct {
	// 数值大小
	value int
	// 叶子节点状态
	available bool
	// 叶子中的排序，方便失效
	rank int
}

func main() {
	var length = 10
	var mm = make(map[int]int, length)
	var o []int

	// 先准备一个顺序随机的数(qie)组(pian)
	for i := 0; i < length; i++ {
		mm[i] = i
	}
	for k, _ := range mm {
		o = append(o, k)
	}

	fmt.Println(o)
	treeSelectionSort(o)
}

func treeSelectionSort(origin []int) []int {
	// 树的层数
	var level int
	var result = make([]int, 0, len(origin))
	for pow(2, level) < len(origin) {
		level++
	}
	// 叶子节点数
	var leaf = pow(2, level)
	var tree = make([]node, leaf*2-1)

	// 先填充叶子节点的数据
	for i := 0; i < len(origin); i++ {
		tree[leaf+i-1] = node{origin[i], true, i}
	}
	// 每层都比较叶子兄弟大小，选出较大值作为父节点
	for i := 0; i < level; i++ {
		// 当前层节点数
		nodeCount := pow(2, level-i)
		// 每组兄弟间比较
		for j := 0; j < nodeCount/2; j++ {
			compareAndUp(&tree, nodeCount-1+j*2)
		}
	}

	// 这个时候树顶端的就是最小的元素了
	result = append(result, tree[0].value)
	fmt.Println(result)

	// 选出最小的元素后，还剩n-1个需要排序
	for t := 0; t < len(origin)-1; t++ {
		// 赢球的节点
		winNode := tree[0].rank + leaf - 1
		// 把赢球的叶子节点状态改为失效
		tree[winNode].available = false

		// 从下一轮开始，只需与每次胜出节点的兄弟节点进行比较
		for i := 0; i < level; i++ {
			leftNode := winNode
			if winNode%2 == 0 {
				leftNode = winNode - 1
			}
			// 比较兄弟节点间大小，并将胜出的节点向上传递
			compareAndUp(&tree, leftNode)
			winNode = (leftNode - 1) / 2
		}
		// 每轮都会吧最小的推到树顶端
		result = append(result, tree[0].value)
		fmt.Println(result)
	}
	return origin
}

func compareAndUp(tree *[]node, leftNode int) {
	rightNode := leftNode + 1

	// 除非左节点无效或者右节点有效并且比左节点大，否则就无脑选左节点
	if !(*tree)[leftNode].available || ((*tree)[rightNode].available && (*tree)[leftNode].value > (*tree)[rightNode].value) {
		(*tree)[(leftNode-1)/2] = (*tree)[rightNode]
	} else {
		(*tree)[(leftNode-1)/2] = (*tree)[leftNode]
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
