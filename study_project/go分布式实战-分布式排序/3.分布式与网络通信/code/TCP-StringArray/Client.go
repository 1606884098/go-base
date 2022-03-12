package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func IntToBytesString(n int) []byte {
	data := int64(n)                                 //数据类型转换
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

func BytesToIntString(bs []byte) int {
	bytebuffer := bytes.NewBuffer(bs) //根据二进制写入二进制结合
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data) //解码
	return int(data)
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端连接失败")
		return
	}
	time.Sleep(time.Second)
	//等待给服务器输入信息
	//00
	//3 4 calc
	//3  8 tasklist
	//01
	go func() {

		myarr := []string{"1", "9", "2", "8", "3", "7", "6", "4", "5", "0", "zxcvb", "asdsd", "bc"}
		//开始
		mybstart := IntToBytesString(0)
		mybstart = append(mybstart, IntToBytesString(0)...)
		conn.Write(mybstart)

		for i := 0; i < len(myarr); i++ {
			mybdata := IntToBytesString(3)
			mybdata = append(mybdata, IntToBytesString(len(myarr[i]))...)
			conn.Write(mybdata)
			conn.Write([]byte(myarr[i]))
		}

		//结束
		mybend := IntToBytesString(0)
		mybend = append(mybend, IntToBytesString(1)...)
		conn.Write(mybend)
	}()
	arr := []string{}
	for {
		//等待，接收信息
		buf := make([]byte, 16)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}

		if n == 16 {
			data1 := BytesToIntString(buf[:len(buf)/2]) //取出第一个数
			data2 := BytesToIntString(buf[len(buf)/2:]) //取出第二个数

			if data1 == 0 && data2 == 0 {
				//开始
				arr = make([]string, 0, 0)
			}
			if data1 == 3 {
				//接收数组
				strbyte := make([]byte, data2, data2)
				length, _ := conn.Read(strbyte)
				fmt.Println(data2, length)
				if length == data2 { //校验长度
					arr = append(arr, string(strbyte))
				}

			}
			if data1 == 0 && data2 == 1 {
				//结束
				fmt.Println("收到数组", arr)
				arr = nil
			}
		}
	}

}
