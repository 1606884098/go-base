package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := divide0(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	//t := divide(2, 0)//捕捉不到
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("程序运行中出现异常：", err)
		}
	}()
	t := divide(2, 0)
	fmt.Println(t)
}
func divide0(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func divide(a, b int) int {
	return a / b
}
