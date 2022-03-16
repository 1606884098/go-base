package TestSlaveServer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func IntToBytes(n int) []byte {
	data := int64(n)                                 //数据类型转换
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

func BytesToInt(bs []byte) int {
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
	//1 23
	//1 32
	//01
	go func() {
		//开始
		mybstart := IntToBytes(0)
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

		myarr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 0}
		for i := 0; i < len(myarr); i++ {
			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(1)...)
			mybdata = append(mybdata, IntToBytes(myarr[i])...)
			conn.Write(mybdata)
		}

		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		mybend = append(mybend, IntToBytes(1)...)
		conn.Write(mybend)
	}()

	arr := []int{}
	for {
		//等待，接收信息
		buf := make([]byte, 24)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}
		fmt.Println(n)
		if n == 24 {
			data1 := BytesToInt(buf[:8])   //取出第一个数
			data2 := BytesToInt(buf[8:16]) //取出第二个数
			data3 := BytesToInt(buf[16:])  //取出第3个数
			fmt.Println(data1, data2, data3)

			if data1 == 0 && data2 == 0 && data3 == 0 {
				//开始
				arr = make([]int, 0, 0)
				fmt.Println("开始")

			}
			if data1 == 0 && data2 == 1 {
				//接收数组
				fmt.Println("收到", data3)
				arr = append(arr, data3)
			}
			if data1 == 0 && data2 == 0 && data3 == 1 {
				//结束
				fmt.Println("收到数组", arr)
				arr = nil
			}
		}
	}

}
