package main

import "fmt"

func main() {
	fmt.Println(rd())
}

func rd() (c, d int) {
	a := 10
	b := 20
	defer func(a, b int) { //必须定义在panic前来处理异常
		err := recover()
		fmt.Println("匿名", a) //变量值为定义在函数钱的值
		fmt.Println("匿名", b)
		fmt.Println("匿名", err)
	}(a, b)
	/*	b=0
		c=a/b*/
	//c=10/0//编译不通过b
	a = 100
	b = 200
	fmt.Println("函数", a)
	fmt.Println("函数", b)
	defer func() {
		fmt.Println("匿名后", a)
		fmt.Println("匿名后", b)
	}()
	return a, b
}
