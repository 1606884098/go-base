package VD02_const01

import (
	"fmt"
	"math"
)

/*
Go语言中的常量使用关键字 const 定义，用于存储不会改变的数据，常量是在编译时被创建的，即
使定义在函数内部也是如此，并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。
由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式。
*/
func main2() {
	/*
		1.常量定义
		1.1标准格式：
		const name [type] = value
		1.2批量声明：
		const (
		 name = value
		 name = value
		)
		所有常量的运算都可以在编译期完成，这样不仅可以减少运行时的工作，也方便其他代码的编译优
		化，当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越
		界、任何导致无效浮点数的操作等。
		常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数
		调用都是返回常量结果：len、cap、real、imag、complex 和 unsafe.Sizeof。
		1.3常量生成器iota
		常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用
		每行都写一遍初始化表达式。在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将
		会被置为 0，然后在每一个有常量声明的行加一。
		1.4无类型的常量
		常量没有一个明确的基础类型。编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少
		有 256bit 的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整
		数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
		1.5字面常量
		所谓字面常量（literal），是指程序中硬编码的常量 如"abc" 1.1 而不是表达式
	*/
	const a int = 1        //1.1
	const c string = "abc" //显式类型定义
	const d = "abc"        //	隐式类型定义

	const ( //1.2
		e  = 2.7182818
		pi = 3.1415926
	)

	type Weekday int //1.3
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	var x float32 = math.Pi //1.4 看看内置函数的实现
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Print(x, y, z)

}
