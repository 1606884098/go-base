package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"time"
)

//整数类型转换成字节
func IntToBytes(n int) []byte {
	data := int64(n)                                 //数据类型转换:为了网络精确都转换成int64
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

//字节转换成整数类型
func BytesToInt(bs []byte) int {
	bytebuffer := bytes.NewBuffer(bs) //根据二进制写入二进制结合
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data) //解码
	return int(data)
}

//浮点类型转换成字节
func Float32ToBytes(data float32) []byte {
	bits := math.Float32bits(data) //math方法，里面没有整数类型方法
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits) //填充
	return bytes
}

//浮点类型转换成字节
func Float64ToBytes(data float64) []byte {
	bits := math.Float64bits(data) //math方法，里面没有整数类型方法
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits) //填充
	return bytes
}

//字节转换成浮点类型
func BytesToFloat32(bs []byte) float32 {
	bits := binary.LittleEndian.Uint32(bs) //解码
	return math.Float32frombits(bits)
}

//字节转换成浮点类型
func BytesToFloat64(bs []byte) float64 {
	bits := binary.LittleEndian.Uint64(bs) //解码
	return math.Float64frombits(bits)
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
		conn.Write(mybstart)
		//表示数据段
		myarr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 0}
		for i := 0; i < len(myarr); i++ {
			mybdata := IntToBytes(1)
			mybdata = append(mybdata, IntToBytes(myarr[i])...)
			conn.Write(mybdata)
		}

		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(1)...)
		conn.Write(mybend)
	}()
	arr := []int{}
	for {
		//等待，接收信息
		buf := make([]byte, 16)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}

		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2]) //取出第一个数
			data2 := BytesToInt(buf[len(buf)/2:]) //取出第二个数

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
				arr = nil
			}
		}
	}

}
