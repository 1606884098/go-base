package main

import "fmt"

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c //这个变量逃逸出来本函数，像这样的情况直接返回b
}

// 空函数, 什么也不做
func void() {
}
func main() {
	// 声明a变量并打印
	var a int
	// 调用void()函数
	void()
	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))

	fmt.Println(dummy1())

}

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy1() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c // 取地址发生逃逸.Go语言最终选择将 c 的 Data 结构分配在堆上。然后由垃圾回收器去回收 c 的内存
}
