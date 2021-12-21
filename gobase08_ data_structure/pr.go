package main

import "fmt"

func main() {
	var a = 10
	p := &a
	*p = 11 //*根据变量的地址取对应的值
	pc := &p
	pd := &pc
	pf := &pd
	fmt.Printf("%p,%p\n", p, pc)
	fmt.Println(*pc)
	fmt.Println(*(*pc))
	fmt.Println(pd)
	fmt.Println(*(*(*(*pf))))

	var pr *int = &a //*int 指针变量，存储的是指针类型，按照int类型的地址操作
	fmt.Println(pr)

}
