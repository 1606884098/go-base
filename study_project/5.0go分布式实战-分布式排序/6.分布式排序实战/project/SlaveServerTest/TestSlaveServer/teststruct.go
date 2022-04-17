package TestSlaveServer

import (
	"fmt"
	"net"
	"time"
)

type Pass struct {
	PassWord string
	times    int
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

		myarr := []Pass{Pass{"abc", 12}, Pass{"axz", 13}, Pass{"zxyz", 10}}
		//开始
		mybstart := IntToBytes(0)
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)
		fmt.Println("send", 0, 0, 0)

		for i := 0; i < len(myarr); i++ {
			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(4)...)
			mybdata = append(mybdata, IntToBytes(myarr[i].times)...)
			conn.Write(mybdata)
			conn.Write(IntToBytes(len(myarr[i].PassWord)))
			conn.Write([]byte(myarr[i].PassWord))
			fmt.Println("send", 0, 0, myarr[i].times, len(myarr[i].PassWord), myarr[i].PassWord)
		}

		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		mybend = append(mybend, IntToBytes(1)...)
		conn.Write(mybend)
		fmt.Println("send", 0, 0, 1)
	}()
	arr := []Pass{}
	for {
		//等待，接收信息
		buf := make([]byte, 24)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器关闭")
			return
		}

		if n == 24 {
			data1 := BytesToInt(buf[:8])   //取出第一个数
			data2 := BytesToInt(buf[8:16]) //取出第二个数
			data3 := BytesToInt(buf[16:])  //取出第3个数
			fmt.Println("收到", data1, data2, data3)

			if data1 == 0 && data2 == 0 && data3 == 0 {
				//开始
				arr = make([]Pass, 0, 0)
			}
			if data1 == 0 && data2 == 4 {
				//接收结构体
				buf1 := make([]byte, 8)
				length, err := conn.Read(buf1)
				fmt.Println("err", err)
				fmt.Println("length", length, string(buf1))
				if length == 8 {
					data4 := BytesToInt(buf1)
					strbyte := make([]byte, data4, data4)
					length, err := conn.Read(strbyte)
					fmt.Println("err", err)
					fmt.Println("strbyte", string(strbyte))
					if length == data4 { //校验长度
						//arr=append(arr,string(strbyte))
						tmpPass := Pass{string(strbyte), data3}
						fmt.Println("收到结构体", tmpPass)
						arr = append(arr, tmpPass)
					}
				}

			}
			if data1 == 0 && data2 == 0 && data3 == 1 {
				//结束
				fmt.Println("收到数组", arr)
				arr = nil
			}
		}
	}

}
