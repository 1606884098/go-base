package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//监听
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		data := make([]byte, 10)
		_, raddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data)
		fmt.Println("Received: ", strData)
		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), raddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Sned: ", upper)
	}

}
