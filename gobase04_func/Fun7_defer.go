package main

import "fmt"

func main() {
	fmt.Println(rd())
}

func rd() (c, d int) {
	a := 10
	b := 20
	defer func(a, b int) { //必须定义在panic前
		err := recover()
		fmt.Println("匿名", a) //变量也是在前定义还要的因为闭包
		fmt.Println("匿名", b)
		fmt.Println("匿名", err)
	}(a, b)
	b = 0
	c = a / b
	//c=10/0//编译不通过b
	a = 100
	b = 200
	fmt.Println("函数", a)
	fmt.Println("函数", b)
	return a, b
}
