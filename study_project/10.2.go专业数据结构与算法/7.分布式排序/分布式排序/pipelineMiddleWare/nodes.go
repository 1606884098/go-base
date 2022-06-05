package pipelineMiddleWare

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var starttime time.Time //构造时间

func Init() {
	starttime = time.Now() //初始化
}

func UseTime() {
	fmt.Println(time.Since(starttime)) //统计消耗时间
}

//内存排序
func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int, 1024) //新的管道
	go func() {
		data := []int{} //创建一个数组，储存数据并且排序
		for v := range in {
			data = append(data, v) //数据压入数组
		}
		fmt.Println("数据读取完成", time.Since(starttime))
		sort.Ints(data) //排序
		for _, v := range data {
			out <- v //压入数据
		}
		fmt.Println("排序完成")
		close(out) //关闭管道
	}()
	return out
}

//合并,两个管道的数据有序，归并有序的数据压入到另外一个管道
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024) //新的管道
	go func() {
		fmt.Println("归并开始")
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		//归并排序，
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1 //取出V1，压入，再次读取v1
				v1, ok1 = <-in1

			} else {
				out <- v2 //取出V2，压入，再次读取v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("归并结束")
	}()

	return out
}

//读取数据
func ReaderSource(reader io.Reader, chunksize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buf := make([]byte, 8) //64
		readsize := 0
		for {
			n, err := reader.Read(buf)
			readsize += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buf)) //数据压入
			}
			if err != nil || (chunksize != -1 && readsize >= chunksize) {
				break //跳出循环
			}
		}

		close(out)
	}()
	return out

}

//写入
func WriterSlink(writer io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)                     //64位 8字节
		binary.BigEndian.PutUint64(buf, uint64(v)) //字节转换
		writer.Write(buf)                          //写入
	}

}

//随机数数组
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int() //压入随机数
		}
		close(out) //关闭管道
	}()

	return out
}

//多路合并5
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	} else {
		m := len(inputs) / 2
		return Merge(MergeN(inputs[:m]...), MergeN(inputs[:m]...)) //递归
	}
}

func ArraySource(num ...int) <-chan int {
	var out = make(chan int)
	go func() {
		for _, v := range num {
			out <- v //数组的数据压入进去
		}
		close(out)
	}()
	return out
}
