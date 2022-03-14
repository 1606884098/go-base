package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func IntToBytes(n int) []byte {
	data := int64(n)                                 //数据类型转换
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

func BytesToInt(bs []byte) int {
	bytebuffer := bytes.NewBuffer(bs) //根据二进制
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

//发送数组
func SendArray(arr []int, conn net.Conn) {
	//开始
	mybstart := IntToBytes(0)
	mybstart = append(mybstart, IntToBytes(0)...)
	mybstart = append(mybstart, IntToBytes(0)...)
	conn.Write(mybstart)

	myarr := arr
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
}

//接收结果
func ServerMsg(conn net.Conn) <-chan int {
	out := make(chan int, 1024)
	defer conn.Close()
	arr := []int{} //接收数据
	buf := make([]byte, 24)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return nil
		}

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
			if data1 == 0 && data2 == 0 && (data3 == 1 || data3 == 2) {
				//结束
				fmt.Println("收到数组", arr)
				for i := 0; i < len(arr); i++ {
					out <- arr[i] //数组压入管道
				}
				close(out)
				return out

				arr = nil
			}
		}
	}

	return nil
}

//归并结果
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 { //可以取出的hua ,就一直取出
			if !ok2 || (ok1 && v1 <= v2) { //轮番取出数据
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("归并排序完成")
	}()

	return out
}

func main() {
	arrlist := [][]int{{1, 100, 2, 93}, {1000, 89, 299, 199}}
	sortResult := []<-chan int{} //结果
	for i := 0; i < 2; i++ {
		tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:"+strconv.Itoa(8848+i))
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpaddr)
		CheckError(err)
		//发送数组
		SendArray(arrlist[i], conn)
		sortResult = append(sortResult, ServerMsg(conn))
		//接收结果
		//归并
	}
	fmt.Println(sortResult)
	fmt.Println(sortResult[0])
	fmt.Println(sortResult[1])
	last := Merge(sortResult[0], sortResult[1]) //归并排序
	fmt.Println("归并以后")
	for v := range last {
		fmt.Printf("%d  ", v)
	}
	time.Sleep(time.Second * 100)

}

func main1() {
	//tcp1  127.0.0.1:8848
	//tcp2  127.0.0.1:8849
	//tcp3  127.0.0.1:8850
	//读取CSDN密码，CSDN密码次数---字符串，结构体
	//切割三份
	//分别发送
	//分别接收
	//waitGroup
	//合并
	//断点续传--容错性

}
