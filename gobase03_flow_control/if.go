package main

import "fmt"

func main() {
	/*	if 条件判断{
		代码语句
	}*/
	if true {
		fmt.Println("条件语句")
	}
	/*	if 条件判断{
			代码语句1
		}else{
			代码语句2
		}*/
	if false {

	} else {

	}
	/*	if 条件判断{
		if 条件判断{
			代码语句1
		}else{
			代码语句2
		}
	}*/
	if true { //建议控制在3层内
		if false {

		} else {

		}
	}

	/*	if 条件判断{
			要执行的代码段
		}else if 条件判断{
			要执行的代码段
		}else if 条件判断{
			要执行的代码段
		}else if条件判断{
			要执行的代码段
		}……………………………else{
		}*/
}
