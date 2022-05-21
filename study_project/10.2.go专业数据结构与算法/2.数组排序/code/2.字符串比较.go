package main

import (
	"fmt"
	"strings"
)

func main1a() {
	fmt.Println(3 > 2)
	//fmt.Println("a" <  "b")//字符串存在地址
	//相等0，不等-1，+1，
	//a<b<c  首先比较第一个字母，左边小于右边 -1 .左边大于右边+1,第一个字母比较不成功比较第二个
	fmt.Println(strings.Compare("ba", "bc"))
	fmt.Println(strings.Compare("c", "b"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a2", "a1"))
	fmt.Println(strings.Compare("a2", "a2"))
	//  aa  ab  ac
}

func main2a() {
	fmt.Println("a" > "b")
	fmt.Println("a" < "b")
	fmt.Println("a" == "a")

	pb := "a1"
	pa := "a2"
	fmt.Println("pa", &pa) //go1.1 ，1.3比较地址，go1.10 优化，比较字符串
	fmt.Println("pb", &pb)
	fmt.Println(pa < pb)
	fmt.Println("a1" < "a2")
	fmt.Println("ab" < "ac")
	fmt.Println("a1b" < "a1c")
}
