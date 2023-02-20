package main

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("init函数1")
}
func init() {
	fmt.Println("init函数3")
}
func init() {
	fmt.Println("init函数2")
}

/*
閉包是不是不利于回收？
defer一般用于数据库连接、文件句柄、锁等
1) 当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer  后的语句压
入到一个栈 中[我为了讲课方便，暂时称该栈为 defer 栈], 然后继续执行函数下一个语句。
2) 当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行(注：遵守栈 先入后出的机制)，
所以同学们看到前面案例输出的顺序。
3) 在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。
*/
func main() {
	f := closure()

	fmt.Println(f()) //通过变量f来调用，f()是执行
	fmt.Println(f())
	c := makeSuffix(".jpg")
	fmt.Println(c("123"))
	fmt.Println(c("456.jpg"))
	fmt.Println(f())
}

/*
闭包是指有权访问另一个函数作用域中的变量的函数，就是在一个函数内部创建另一个函数。
虽然不能在一个函数里直接声明另一个函数，但是可以在一个函数中声明一个函数类型的变
量，此时的函数称为闭包（closure),所有的匿名函数(Go语言规范中称之为函数字面量)都是闭包。

它不关心这些捕获了的变量和常量是否已经超出了作用域，所以只有闭包还在使用它，这些变量就
还会存在。
*/
func closure() func() int { //累加器
	var x int
	return func() int {
		x++
		return x
	}
}

/*
1) 编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，
并返回一个闭包
2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回
文件名.jpg , 如果已经有.jpg 后缀，则返回原文件名。
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果没有后缀加上，有后缀直接返回
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
