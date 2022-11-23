package main

import (
	"fmt"
	"net"
	_ "os/exec"
)

/*
Go语言是静态类型语言，因此变量（variable）是有明确类型的，编译器也会检查变量类型的正确性。在数学概念中，变量表示没有固定值且可改变的数。但从
计算机系统实现角度来看，变量是一段或多段用来存储数据的内存。

Go语言的词法元素包括 5 种，分别是标识符（identifier）、关键字（keyword）、操作符（operator）、分隔符（delimiter）、字面量（literal），它们是
组成Go语言代码和程序的最基本单位。

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
		name：是变量名（标识符，并且严格区分大小写,首字母大写为共有，小写为私有）
		type：是变量的类型
		关于标识符：
		由 26 个英文字母、0~9、_组成；
		不能以数字开头，例如 var 1num int 是错误的；
		Go语言中严格区分大小写；
		标识符不能包含空格；
		不能以系统保留关键字作为标识符。
		关键字（25）：break default func interface select
		case defer go map struct
		chan else goto package switch
		const fallthrough if range type
		continue for import return var
		预定义标识符：append bool byte cap close complex complex64 complex128 uint16
		copy false float32 float64 imag int int8 int16 uint32
		int32 int64 iota len make new nil panic uint64
		print println real recover string true uint uint8 uintptr

		命名标识符时还需要注意以下几点：
		标识符的命名要尽量采取简短且有意义；
		不能和标准库中的包名重复；
		为变量、函数、常量命名时采用驼峰命名法，

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
	var a int //1.1

	var ( //1.2
		b string
		c []float32
		d func() bool
		e struct {
			f int
		}
	)

	g, h := 0, 1 //1.3
	//go语言中定义的变量一定要使用，定义不使用编译报错。当一个变量被声明之后，系统自动赋予它该类型的零值：
	// int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。所有的内存在 Go 中都是经过初始化的
	fmt.Print(a, b, c, d, e, g, h)
	/*
		2.变量初始化
		2.1变量初始化的标准格式
		var name type = express

		2.2编译器推导类型的格式
		var name=express
		编译器会尝试根据等号右边的表达式推导name变量的类型,等号右边的部分在编译原理里被称做右值（rvalue）。

		2.3短变量声明并初始化
		name_1,name_2:=express_1,express_2
	*/
	var i int = 10    //2.1 或 var i int   i=10
	var j = "str"     //2.2
	k, l := 10, "str" //2.3
	fmt.Print(i, j, k, l)
	/*
		3.左值与右值：等号右边的部分在编译原理里被称做右值（rvalue），左值就是可以改变的值
	*/
	var hp = 100 //3 hp是左值，可以改变的值
	//100=129//不能改变，只能赋值给左值的值
	fmt.Print(hp)
	/*
		4.变量赋值
		变量可以重复赋值,一旦给一个变量赋了新值,那么变量中的老值就不复存在了
		一个变量赋值是从右边往左边；多重赋值时，变量的左值和右值按从左到右的顺序赋值。

		= 和:=的区别：
		=是给已有的变量赋值,已经分配好了内存空间
		:=分配内存空间并且赋值，自动推导类型赋值

		自动推导类型:就是不用通过var 声明变量，不用指定类型，直接在变量名后面跟”:”号，同时完成赋值。那么GO会根据所赋的值自动
		推导出变量的类型。
	*/
	//多重赋值交换变量值
	var m int = 100
	var n int = 200
	n, m = m, n

	/*为了节省内存也可以这么玩
	m=m^n
	n=n^m
	m=m^n
	*/
	fmt.Println(m, n)

	/*
		5.匿名变量:
		在编码过程中，可能会遇到没有名称的变量、类型或方法。虽然这不是必须的，但有时候这样做可以极大地增强代码的灵活性，这些变量被统称为匿
		名变量。

		5.1"_"接受匿名变量，相当于是占位符
		_ 匿名变量，丢弃数据不进行处理, _匿名变量配合函数返回值使用才有价值。
		匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

		5.2"_"导包的时候使用
		import 下划线的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进
		来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以
		无法通过包名来调用包中的其他函数。
	*/
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn)

	var a1 int = 10
	var a2 int = 20
	swap(&a1, &a2)
	fmt.Printf("a1=%d,a2=%d\n", a1, a2)
}

func swap(i *int, i2 *int) {
	*i, *i2 = *i2, *i
	fmt.Printf("i=%d,i2=%d\n", *i, *i2)
}
