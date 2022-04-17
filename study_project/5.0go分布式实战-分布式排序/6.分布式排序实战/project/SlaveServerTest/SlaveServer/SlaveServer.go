package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type Pass struct {
	PassWord string
	times    int
}

func Server(conn net.Conn) {
	if conn == nil {
		fmt.Println("无效连接")
		return
	}
	//接收数据，处理
	arr := []interface{}{} //泛用
	mytype := 0            //标记类型

	filepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\tmpfile\\"
AAA:
	rand.Seed(time.Now().UnixNano())
	filenamenumber := rand.Int() % 10000
	filestorepath := filepath + strconv.Itoa(filenamenumber) + ".txt"
	_, err := os.Stat(filestorepath) //判断文件
	if err == nil {
		goto AAA
	}
	fmt.Println(filestorepath)
	savefile, _ := os.Create(filestorepath)
	var save *bufio.Writer
	save = bufio.NewWriter(savefile) //保存

	for {
		//等待，接收信息
		buf := make([]byte, 24)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端关闭")
			return
		}
		fmt.Println("n=", n)
		if n == 24 {
			data1 := BytesToInt(buf[:8])   //取出第一个数
			data2 := BytesToInt(buf[8:16]) //取出第二个数
			data3 := BytesToInt(buf[16:])  //取出第3个数
			fmt.Println("receive", data1, data2, data3)

			if data1 == 0 {
				//内存模式
				if data2 == 0 && data3 == 0 {
					//开始
					arr = make([]interface{}, 0, 0)

				}
				if data2 == 1 {
					//整数
					arr = append(arr, data3)
					mytype = 1

				} else if data2 == 2 {
					//实数
					arr = append(arr, ByteToFloat64(buf[16:]))
					mytype = 2

				} else if data2 == 3 {
					//字符串
					strbyte := make([]byte, data3, data3)
					length, _ := conn.Read(strbyte)
					if length == data3 { //校验长度
						arr = append(arr, string(strbyte))
					}
					mytype = 3

				} else if data2 == 4 {
					//结构体
					fmt.Println("收到结构体指示")
					buf1 := make([]byte, 8)
					length, _ := conn.Read(buf1)
					fmt.Println("length", length, string(buf1))
					if length == 8 {
						data4 := BytesToInt(buf1)
						strbyte := make([]byte, data4, data4)
						length, _ := conn.Read(strbyte)
						fmt.Println("strbyte", string(strbyte), "length", length)
						if length == data4 { //校验长度
							//arr=append(arr,string(strbyte))
							tmpPass := Pass{string(strbyte), data3}
							fmt.Println("收到结构体", tmpPass)
							arr = append(arr, tmpPass)
						}
					}
					mytype = 4

				}

				if data2 == 0 && data3 == 1 {
					//结束，从小到大
					fmt.Println("收到数组", arr)
					//排序 从小到大
					mydata := new(QuickSortData)
					mydata.Data = arr
					mydata.IsSmalltoBig = true
					QuickSortDataByType(mydata, mytype)
					mydata.QuickSort()
					fmt.Println("最终数组", arr)
					arr = nil
					runtime.GC()
					debug.FreeOSMemory() //释放内存
					SendResult(mydata, conn, mytype, true)
				}
				if data2 == 0 && data3 == 2 {
					//结束，从大到小
					fmt.Println("收到数组", arr)
					//排序 从大到小
					mydata := new(QuickSortData)
					mydata.Data = arr
					mydata.IsSmalltoBig = false
					QuickSortDataByType(mydata, mytype)
					mydata.QuickSort()

					fmt.Println("最终数组", arr)
					arr = nil
					runtime.GC()
					debug.FreeOSMemory() //释放内存
					SendResult(mydata, conn, mytype, false)

				}
			} else if data1 == 1 {
				//硬盘模式

				if data2 == 0 && data3 == 0 {
					//开始
					fmt.Println("开始")

				}
				if data2 == 1 {
					//整数
					fmt.Fprintln(save, strconv.Itoa(data3))
					//arr=append(arr,data3)
					mytype = 1

				} else if data2 == 2 {
					//实数
					floatnum := ByteToFloat64(buf[16:])
					fmt.Fprintln(save, strconv.FormatFloat(floatnum, 'f', 6, 64))
					//arr=append(arr,data3)
					mytype = 2
				} else if data2 == 3 {
					//字符串

					strbyte := make([]byte, data3, data3)
					length, _ := conn.Read(strbyte)
					if length == data3 { //校验长度
						//arr=append(arr,string(strbyte))
						fmt.Fprintln(save, string(strbyte))
					}
					mytype = 3

					//arr=append(arr,data3)

				} else if data2 == 4 {
					//结构体
					//fmt.Fprintln(save,strconv.Itoa(data3))
					//arr=append(arr,data3)
					buf1 := make([]byte, 8)
					length, _ := conn.Read(buf)
					if length == 8 {
						data4 := BytesToInt(buf1)
						strbyte := make([]byte, data4, data4)
						length, _ := conn.Read(strbyte)
						if length == data4 { //校验长度
							//arr=append(arr,string(strbyte))
							//tmpPass:=Pass{string(strbyte),data3}
							fmt.Fprintln(save, string(strbyte)+" # "+strconv.Itoa(data3))
							//arr=append(arr,tmpPass)
						}
					}
					mytype = 4
				}

				if data2 == 0 && data3 == 1 {
					//结束，从小到大
					fmt.Println("结束", 1)
					save.Flush()
					savefile.Close()
					//类型划分
					//读取文件
					//排序
					//写入文件
					//读取文件传输

					DiskfileSortAndSend(filestorepath, mytype, true, conn)

				}
				if data2 == 0 && data3 == 2 {
					fmt.Println("结束", 2)
					//结束，从大到小
					save.Flush()
					savefile.Close()

					DiskfileSortAndSend(filestorepath, mytype, false, conn)

				}

			}

		}

	}
}

func DiskfileSortAndSend(filepath string, mytype int, isSmalltoBig bool, conn net.Conn) {
	arr := []interface{}{} //数组

	fi, err := os.Open(filepath)
	defer fi.Close() //打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine() //读取行数
		if err == io.EOF {
			break //文件结束，末尾
		}
		if mytype == 1 {
			data, _ := strconv.Atoi(string(line))
			arr = append(arr, data)
		} else if mytype == 2 {
			data, _ := strconv.ParseFloat(string(line), 64)
			arr = append(arr, data)
		} else if mytype == 3 {
			data := string(line)
			arr = append(arr, data)
		} else if mytype == 4 {
			data := string(line)
			datalist := strings.Split(data, " # ")
			times, _ := strconv.Atoi(datalist[1])
			arr = append(arr, Pass{datalist[0], times})
		}
	}
	//排序
	mydata := new(QuickSortData)
	mydata.IsSmalltoBig = true
	mydata.Data = arr
	QuickSortDataByType(mydata, mytype)
	mydata.QuickSort()
	fmt.Println("最终数组", arr)

	newpath := strings.Replace(filepath, ".txt", "sort.txt", -1)
	savefile, _ := os.Create(newpath)
	save := bufio.NewWriter(savefile) //保存
	for i := 0; i < len(mydata.Data); i++ {
		//fmt.Fprintln(save,strconv.Itoa(mydata.Data[i].(int)))
		if mytype == 1 {
			fmt.Fprintln(save, strconv.Itoa(mydata.Data[i].(int)))
		} else if mytype == 2 {
			fmt.Fprintln(save, strconv.FormatFloat(mydata.Data[i].(float64), 'f', 6, 64))
		} else if mytype == 3 {
			fmt.Fprintln(save, mydata.Data[i].(string))
		} else if mytype == 4 {
			fmt.Fprintln(save, mydata.Data[i].(Pass).PassWord+" # "+strconv.Itoa(mydata.Data[i].(Pass).times))
		}
	}
	save.Flush()
	savefile.Close()

	//释放内存
	arr = nil
	runtime.GC()
	debug.FreeOSMemory() //释放内存

	//发送
	//整数
	mybstart := IntToBytes(1) //内存
	mybstart = append(mybstart, IntToBytes(0)...)
	mybstart = append(mybstart, IntToBytes(0)...)
	conn.Write(mybstart)
	fmt.Println("send", 1, 0, 0)

	fisort, err := os.Open(newpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	brsort := bufio.NewReader(fisort)
	for {
		line, _, err := brsort.ReadLine() //读取行数
		if err == io.EOF {
			break //文件结束，末尾
		}
		if mytype == 1 {
			data, _ := strconv.Atoi(string(line))
			//arr=append(arr,	data)
			//conn.Write(IntToBytes(1))
			//conn.Write(IntToBytes(1))
			//conn.Write(IntToBytes(data))
			mybdata := IntToBytes(1)
			mybdata = append(mybdata, IntToBytes(1)...)
			mybdata = append(mybdata, IntToBytes(data)...)
			conn.Write(mybdata)

		} else if mytype == 2 {
			data, _ := strconv.ParseFloat(string(line), 64)
			//arr=append(arr,	data)
			//conn.Write(IntToBytes(1))
			//conn.Write(IntToBytes(2))
			//conn.Write(Float64ToByte(data))
			mybdata := IntToBytes(1)
			mybdata = append(mybdata, IntToBytes(2)...)
			mybdata = append(mybdata, Float64ToByte(data)...)
			conn.Write(mybdata)

		} else if mytype == 3 {
			//conn.Write(IntToBytes(1))
			//conn.Write(IntToBytes(3))
			//mybdata:=IntToBytes(len(string(line)))
			//conn.Write(mybdata)
			mybdata := IntToBytes(1)
			mybdata = append(mybdata, IntToBytes(3)...)
			mybdata = append(mybdata, IntToBytes(len(string(line)))...)
			conn.Write(mybdata)
			conn.Write(line)

		} else if mytype == 4 {
			data := string(line)
			datalist := strings.Split(data, " # ")
			password := datalist[0]
			times, _ := strconv.Atoi(datalist[1])

			//arr=append(arr,	data)
			//conn.Write(IntToBytes(1))
			//conn.Write(IntToBytes(4))
			//conn.Write(IntToBytes(times))
			mybdata := IntToBytes(1)
			mybdata = append(mybdata, IntToBytes(4)...)
			mybdata = append(mybdata, IntToBytes(times)...)
			conn.Write(mybdata)

			conn.Write(IntToBytes(len(password)))
			//104
			conn.Write([]byte(password))

		}

	}
	fi.Close() //打开文件

	//结束
	mybend := IntToBytes(1)
	mybend = append(mybend, IntToBytes(0)...)
	if isSmalltoBig {
		mybend = append(mybend, IntToBytes(1)...)
		fmt.Println("send", 1, 0, 1)
	} else {
		mybend = append(mybend, IntToBytes(2)...)
		fmt.Println("send", 1, 0, 2)
	}
	conn.Write(mybend)

}

func SendResult(mydata *QuickSortData, conn net.Conn, mytype int, isSmalltoBig bool) {
	if mytype == 1 {
		//整数
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

		fmt.Println("send", 0, 0, 0)

		for i := 0; i < len(mydata.Data); i++ {

			//conn.Write(IntToBytes(0))
			//mybdata:=IntToBytes(1)
			//conn.Write(mybdata)
			//conn.Write(IntToBytes(mydata.Data[i].(int)))
			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(1)...)
			mybdata = append(mybdata, IntToBytes(mydata.Data[i].(int))...)
			conn.Write(mybdata)
			fmt.Println("send", 0, 1, mydata.Data[i].(int))
		}
		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
			fmt.Println("send", 0, 0, 1)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
			fmt.Println("send", 0, 0, 2)
		}
		conn.Write(mybend)

	} else if mytype == 2 {
		//实数
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

		for i := 0; i < len(mydata.Data); i++ {

			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(2)...)
			mybdata = append(mybdata, Float64ToByte(mydata.Data[i].(float64))...)
			conn.Write(mybdata)

		}
		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
		}
		conn.Write(mybend)

	} else if mytype == 3 {
		//字符串
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)
		fmt.Println("send", 0, 0, 0)

		for i := 0; i < len(mydata.Data); i++ {
			//conn.Write(IntToBytes(0))
			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(3)...)
			mybdata = append(mybdata, IntToBytes(len(mydata.Data[i].(string)))...)

			conn.Write(mybdata)
			conn.Write([]byte(mydata.Data[i].(string)))
			fmt.Println("send", 0, 3, len(mydata.Data[i].(string)), mydata.Data[i].(string))
		}
		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
			fmt.Println("send", 0, 0, 1)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
			fmt.Println("send", 0, 0, 2)
		}
		conn.Write(mybend)

	} else if mytype == 4 {
		//结构体
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)
		fmt.Println("send", 0, 0, 0)

		for i := 0; i < len(mydata.Data); i++ {
			//conn.Write(IntToBytes(0))
			mybdata := IntToBytes(0)
			mybdata = append(mybdata, IntToBytes(4)...)
			mybdata = append(mybdata, IntToBytes(mydata.Data[i].(Pass).times)...)
			conn.Write(mybdata)
			conn.Write(IntToBytes(len(mydata.Data[i].(Pass).PassWord)))
			conn.Write([]byte(mydata.Data[i].(Pass).PassWord))
			fmt.Println("send", 0, 4, mydata.Data[i].(Pass).times, len(mydata.Data[i].(Pass).PassWord), mydata.Data[i].(Pass).PassWord)
		}
		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
		}
		conn.Write(mybend)

	}
}

func QuickSortDataByType(mydata *QuickSortData, mytype int) {
	if mytype == 1 {
		mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
			if IsSmalltoBig {
				return data1.(int) < data2.(int)
			} else {
				return data1.(int) > data2.(int)
			}
		}
	} else if mytype == 2 {
		mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
			if IsSmalltoBig {
				return data1.(float64) < data2.(float64)
			} else {
				return data1.(float64) > data2.(float64)
			}
		}
	} else if mytype == 3 {
		mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
			if IsSmalltoBig {
				return data1.(string) < data2.(string)
			} else {
				return data1.(string) > data2.(string)
			}
		}
	} else if mytype == 4 {
		mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
			if IsSmalltoBig {
				return data1.(Pass).times < data2.(Pass).times
			} else {
				return data1.(Pass).times > data2.(Pass).times
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
	fmt.Println("服务器开启成功")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("连接出错")
		}
		go Server(conn) //并发处理
	}

}
