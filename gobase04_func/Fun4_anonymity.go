package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
func main() {
	// 1.匿名函数做回调函数使用，使用匿名函数打印切片内容
	visit([]int{1, 4, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	flag.Parse()
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
		fmt.Printf("%T\n", f)
		f()
	} else {
		fmt.Println("skill not found")
	}

}
