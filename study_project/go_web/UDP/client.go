package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}

	defer conn.Close()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//接收用户输入
		line := input.Text()
		length := len(line)
		n := 0
		for i := 0; i < length; i += n {
			var toWrite string
			if length-i > 10 {
				toWrite = line[i : i+10]
			} else {
				toWrite = line[i:]
			}
			//将数据写入连接
			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			msg := make([]byte, 10)
			_, err = conn.Read(msg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("收到:", string(msg))
		}
	}
}
