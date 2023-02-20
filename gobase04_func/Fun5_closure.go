package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is fi")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

/*
- 闭包：由函数以及相关引用环境组合而成的实例，也就是说``闭包=函数+引用环境``
- 匿名函数：匿名函数就是我们说的闭包，它不能独立存在，但可以直接调用或者赋值于某个变量。
一个闭包，继承了函数声明时的作用域。在go语言中，所有的匿名函数都是闭包
*/
func main() {
	ret := f3(f2, 100, 200) // 把原来需要传递两个int类型的参数包装成一个不需要传参的区数
	fmt.Printf("%T\n", ret)
	f1(ret)

}
