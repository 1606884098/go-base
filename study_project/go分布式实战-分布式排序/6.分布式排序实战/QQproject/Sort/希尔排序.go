package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ShellSortGo(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	cpunum := runtime.NumCPU() //获取CPU数量
	wg := sync.WaitGroup{}     //批量等待

	//压缩空间
	for gap := len(arr); gap > 0; gap /= 2 {
		wg.Add(cpunum)
		ch := make(chan int, 10000)
		go func() {
			//管道写入
			for k := 0; k < gap; k++ {
				ch <- k
			}
			close(ch) //关闭管道
		}()
		for k := 0; k < cpunum; k++ {
			go func() {
				for v := range ch {
					ShellSortStep(arr, v, gap)
				}
				wg.Done()
			}()

		}

		wg.Wait()

	}

	fmt.Println(arr)

}

// 1  9   2  8  3  7  4   6  5 10  13
//1                7                13
//   9                4
//       2                6
//          8                  5
//              3                  10

//1                7
//   4                9
//       2                6
//          5                 9
//              3                  10

//1  4   2  5  3   7  9 6    9  10
//1         5         9         10
//   4        3         6
//       2         7         9

//1         5         9         10
//   3        4         6
//       2         7         9

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { //循环
		backup := arr[i] //备份插入的数据
		j := i - gap     //上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] //从后往前移动
			j -= gap            //跳过步长
		}
		arr[j+gap] = backup //插入备份
		fmt.Println(arr)

	}

}

func ShellSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				//每个元素单独处理步长
				ShellSortStep(arr, i, gap)
			}

			gap = gap / 2
		}

	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 112, 102}
	fmt.Println(arr)
	ShellSortGo(arr)
	fmt.Println(arr)
}
