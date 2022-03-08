package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sort"
)

func IntToBytesArr(n int) []byte {
	data := int64(n)                                 //数据类型转换
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

func BytesToIntArr(bs []byte) int {
	bytebuffer := bytes.NewBuffer(bs) //根据二进制写入二进制结合
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data) //解码
	return int(data)
}

func Server(conn net.Conn) {
	if conn == nil {
		fmt.Println("无效连接")
		return
	}
	//接收数据，处理
	arr := []int{}

	for {
		//等待，接收信息
		buf := make([]byte, 16)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端关闭")
			return
		}
		if n == 16 {
			data1 := BytesToIntArr(buf[:len(buf)/2]) //取出第一个数
			data2 := BytesToIntArr(buf[len(buf)/2:]) //取出第二个数

			if data1 == 0 && data2 == 0 {
				//开始
				arr = make([]int, 0, 0)
			}
			if data1 == 1 {
				//接收数组
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1 {
				//结束
				fmt.Println("收到数组", arr)
				sort.Ints(arr) //排序
				fmt.Println("数组排序完成", arr)

				//返回结果
				myarr := arr

				mybstart := IntToBytesArr(0)
				mybstart = append(mybstart, IntToBytesArr(0)...)
				conn.Write(mybstart)

				for i := 0; i < len(myarr); i++ {
					mybdata := IntToBytes(1)
					mybdata = append(mybdata, IntToBytesArr(myarr[i])...)
					conn.Write(mybdata)
				}

				//结束
				mybend := IntToBytes(0)
				mybend = append(mybend, IntToBytesArr(1)...)
				conn.Write(mybend)

			}
		}

	}
}

func main() {

	server, err := net.Listen("tcp", "127.0.0.1:8848")
	defer server.Close()
	if err != nil {
		fmt.Println("服务器开启失败")
		return
	}
	fmt.Println("正在开启服务器....")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("连接出错")
		}
		go Server(conn) //并发处理
	}

}
