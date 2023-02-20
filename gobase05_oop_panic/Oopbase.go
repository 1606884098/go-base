package main

import (
	"fmt"
)

type person struct {
	id   int
	name string
	age  int
}

func (p *person) talk() {
	fmt.Println("说英语！")
}

type student struct {
	*person //也可以是指针
	//person
	name  string
	score float32
}

//中国学生说中文，只能重写方法
func (s *student) talk() {
	fmt.Println("说中文！")
}

/*多态参数
在前面的 Usb 接口案例，Usb usb ，即可以接收手机变量，又可以接收相机变量，
就体现了 Usb  接 口 多态。
多态数组
演示一个案例：给 Usb 数组中，存放 Phone 结构体 和	Camera 结构体变量 案例说明:*/
//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让 Phone 实现 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	name string
}

//让 Camera 实现	Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	/*
		类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值
		是否实现了期望的接口或者具体的类型。
		value, ok := x.(T)
		其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。
		该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x
		是否为 T 类型
	*/
	var x Usb
	y, ok := x.(Camera)
	fmt.Println(y, ok)

	var s1 student
	s1.name = "好人"
	s1 = student{&person{1, "国家", 1}, "", 98}
	s1.talk()        //重写的方法
	s1.person.talk() //直接调用方法

	//定义一个 Usb 接口数组，可以存放 Phone 和 Camera 的结构体变量
	//这里就体现出多态数组，主要用于入参可以接受多个参数，从而提升扩展性
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	var cc Usb
	cc = usbArr[0]
	cc.Start()

	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
}
