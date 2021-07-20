package main

import "fmt"

type Person struct {
	name string
}

func (p Person) testMathod() { //结构体绑定方法
	fmt.Println(p.name)
}

func (p *Person) testmathod1() {
	fmt.Println("v=", p.name)
}

func main() {
	var p Person
	p.name = "jack"
	p.testMathod() //调用方法
	p.testmathod1()
}
