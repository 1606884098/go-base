package main

import "fmt"

/*一个文件可以有多个init函数*/
func init() {
	fmt.Println("另一个文件init-1")
}
func init() {
	fmt.Println("另一个文件init-2")
}
func test() {
	fmt.Println("测试init函数")
}
