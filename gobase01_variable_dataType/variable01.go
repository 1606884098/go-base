package main

import "fmt"

/*
Go语言是静态类型语言，因此变量（variable）是有明确类型的，编译器也会检查变量类型的正确性。在数学概念中，变量表示没有固定值且可改变的数。但从
计算机系统实现角度来看，变量是一段或多段用来存储数据的内存。

静态语言和动态语言的区别：
静态语言：
变量的类型在编译之前就需要确定，在编译的时候需要先编译，将源码转换成目标代码，然后需要运行目标代码程序才能运行，比如go,C++、Java、Delphi、C#。
动态语言：
不需要直接指定变量类型，在解释的时候，转换为目标代码和运行程序一步到位，比如Python、Ruby、Perl.可以在运行时改变结构.
*/
func main() {
	/*
		1.定义变量
		1.1 标准格式：var name type
		var：是声明变量的关键字
		name：是变量名（标识符：有非数字开头，字母，数字，下划线组成）
		type：是变量的类型
		1.2 批量格式
			var(
				name_1 type
				name_2 type
			)
		var 形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。
		1.3 简短格式
		name:= express 或 name_1,name_2:= express_1,express_2
		需要注意的是，简短模式（short variable declaration）有以下限制：
		定义变量，同时显式初始化。
		不能提供数据类型。
		只能用在函数内部。
		简短变量声明被广泛用于大部分的局部变量的声明和初始化
	*/
	var a int

	var (
		b string
		c []float32
		d func() bool
		e struct {
			f int
		}
	)

	i, j := 0, 1
	//go语言中定义的变量一定要使用，定义不使用编译报错
	fmt.Print(a, b, c, d, e, i, j)
	/*
		2.变量初始化
		2.1变量初始化的标准格式
		var name type = express
		2.2编译器推导类型的格式
		var name=express
		编译器会尝试根据等号右边的表达式推导 hp 变量的类型,等号右边的部分在编译原理里被称做右值（rvalue）。
		2.3短变量声明并初始化
	*/
}
