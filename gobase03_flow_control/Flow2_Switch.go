package main

import "fmt"

func main() {
	/*	switch 变量或者表达式的值{
		case 值1:
			要执行的代码
		case 值2:
			要执行的代码
		case 值3:
			要执行的代码
			………………………………..
		default:
			要执行的代码
		}*/
	switchTest(2) //switch大部分是值
	ifElseTest(2) //if else可以是范围
}

func ifElseTest(i int) {
	if i > 1 {
		fmt.Println("i>1")
	} else if i > 1 && i < 4 {
		fmt.Println("i>1&&i<4")
	} else if i > 4 && i < 6 {
		fmt.Println("i>4&&i<6")
	} else if i > 6 && i < 8 {
		fmt.Println("i>6&&i<8")
	} else {
		fmt.Println("你搞错了")
	}
}

func switchTest(i int) {
	var l = i
	switch l {
	case 1, 5, 6:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
}
