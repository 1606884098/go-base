package main

import (
	"fmt"
	"unsafe"
)

/*
//slice是一个把go数组进行了包装的一个结构体，但是这个结构体只是在编译等其他层面能看到
// 如果想在运行期间使用的话可以使用其对应的reflect结构体
// 即reflect.SliceHeader
type slice struct {src/runtime/slice
array unsafe.Pointer //指向存放数据的数组指针
len int //长度有多大
cap int //容量有多大
}
*/
func main() {

	/*var 变量名 []类型*/
	var dd []int
	if dd == nil {
		fmt.Printf("%p\n", dd) //没有分配内存0x0
	}

	var a []int = []int{1, 2, 3} //直接定义
	ptr := unsafe.Pointer(&a[0])
	fmt.Printf("%p\n,%p\n", a, ptr) //分配内存0xc000010380
	b := []int{1, 2, 3}             //直接定义
	fmt.Printf("slice:%v,len=%v,cap=%v\n", b, len(b), cap(b))
	var c []int = make([]int, 5) //长度(len)是5，容量(cap)也是5（这里的容量是值初始容量如果添加原始会自动扩容）

	var d []int = make([]int, 5, 10) //长度(len)是5，容量(cap)是10

	arr := [5]int{1, 2, 3, 4, 5} //通过数组定义
	var e []int
	e = arr[1:4] // 前包后不包
	fmt.Println(a, b, c, d, e)
	/*初始化
	全局：
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[start:end]
	var slice1 []int = arr[:end]
	var slice2 []int = arr[start:]
	var slice3 []int = arr[:]
	var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
	局部：
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[start:end]
	slice6 := arr[:end]
	slice7 := arr[start:]
	slice8 := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素*/

	/*切片的本质是一个指针，指向一个数据 如下例子*/
	fmt.Printf("切片e的值：%v\n", e)
	fmt.Printf("数组的值：%v\n", arr)
	e[0] += 10
	e[1] += 10
	p := &e[2]
	*p = 199 //通过指针修改数据
	fmt.Printf("切片修改后的值e的值：%v\n", e)
	fmt.Printf("数组修改后的值的值：%v\n", arr)

	data2 := [][]int{[]int{1, 2}, []int{1, 2, 3}, []int{6, 2, 3}} //二维切片
	fmt.Printf("二维切片：%v\n", data2)
	data3 := [][][]int{[][]int{[]int{1, 2}, []int{1, 2, 3}, []int{6, 2, 3}},
		[][]int{[]int{1, 2}, []int{1, 2, 3}, []int{6, 2, 3}}} //三维切片
	fmt.Printf("三维切片：%v\n", data3)

	/*扩容原理，不够就翻倍*/
	var data4 = []int{1, 2, 3}
	fmt.Println(data4, len(data4), cap(data4))
	for i := 0; i < 10; i++ {
		data4 = append(data4, i+10) //追加元素，内存不够重新分配内存,地址变化，容量翻倍
		fmt.Printf("%p", data4)
		fmt.Println(data4, len(data4), cap(data4))
	}

	data5 := data4[0:2:4] //后面的数字4是容量
	fmt.Println(data5, len(data5), cap(data5))
	fmt.Printf("%p", data5)
	fmt.Println(data5, len(data5), cap(data5))
	data5 = append(data5, 10, 23, 23, 24, 31, 41) //超原来slice.cap底层就会重新分配内存，即使原数组没填满
	fmt.Printf("%p", data5)
	fmt.Println(data5, len(data5), cap(data5))

	/*string底层就是一个byte的数组，因此，也可以进行切片操作。*/
	str := "dsafds中国"
	str1 := str[6:]
	fmt.Printf(str1)

	/*切片遍历*/
	for i, v := range data5 {
		fmt.Printf("index=%d,val=%d\n", i, v)
	}
}
