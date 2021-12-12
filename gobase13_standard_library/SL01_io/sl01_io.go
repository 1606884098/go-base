package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))*/

	sc := bufio.NewScanner(os.Stdin) //扫描键盘的输入
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

}
