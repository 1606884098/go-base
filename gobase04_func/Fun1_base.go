package main

import (
	"flag"
	"fmt"
)

/*
func 函数名(形式参数列表)(返回值列表){
    函数体
}
函数执行到代码块最后一行}之前或者 return
语句的时候会退出，其中 return 语句可以带有零个或多个参数
函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的
变量作用域：尽量使用局部变量栈上分配
(1)在函数外边定义的变量叫做全局变量，全局变量能够在所有的函数中进行访问
(2) 如果全局变量的名字和局部变量的名字相同，那么使用的是局部变量的，小技巧强龙不压地头蛇
*/
func main() {
	a, b, _ := namedRetValues(1, 2) //(1,2)实参，要与形参的个数和类型保持一致
	println(a, b)
	MyPrintf(1, "q", 'q')
	ss := anonymity("匿名函数", func(i int) bool {

		return false
	})
	println("ss" + ss)
}

//形参(int,int)
func namedRetValues(int, int) (int, int, int) { //(a, b int, int)在函数参数中混合使用了命名和非命名参数报错
	for i := 0; i < 10; i++ {
		println(1)
		return 1, 2, 3
	}
	return 0, 0, 0
}

/*任意类型的可变参数也可以是指定类型 混合用可变参数(a int,args ...string{}) 但是
可变参数一定放在最后

两种传递方式
	1) 值传递
	2) 引用传递 其实，不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，
	值传递的是值的拷贝，引用传递的是地址的拷贝，一般来说，地址拷贝效率高，因为数据量小
	，而值拷贝决定拷贝的 数据大小，数据越大，效率越低。
*/
func MyPrintf(a int, args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

/*
匿名函数：
f:=func(参数列表)(返回参数列表){
    函数体
}(参数列表1)
(参数列表1)表示对匿名函数进行调用，传递参数为(参数列表1),再赋值给f
匿名函数做参数，匿名函数赋值，也可以直接使用
*/
func anonymity(s string, f func(int) bool) string {
	f1 := func(data int) int {
		fmt.Println("hello", data)
		return data
	} //(100)
	// 使用f()调用
	f1(100)
	println("cc:", f1)
	//return "zheng"

	var skillParam = flag.String("skill", "", "skill to perform")
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
	return "zheng"
}
