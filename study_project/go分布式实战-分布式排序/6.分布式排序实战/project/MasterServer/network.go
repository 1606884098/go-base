package main

import "net"
import "os"
import "bufio"
import "fmt"
import "io"
import "strconv"
import "strings"

//发送数据
func SendFile(sendfilepath string, conn net.Conn, datatype int, isMEM bool, SmalltoBig bool) {
	if isMEM {
		//控制slaveserver

		mybstart := IntToBytes(0)
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

		//读取文件制定类型发送之,
		fi, err := os.Open(sendfilepath)
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
			mytype := datatype
			if mytype == 1 {
				data, _ := strconv.Atoi(string(line))
				//发送整数
				mybdata := IntToBytes(0)
				mybdata = append(mybdata, IntToBytes(1)...)
				mybdata = append(mybdata, IntToBytes(data)...)
				conn.Write(mybdata)

			} else if mytype == 2 {
				data, _ := strconv.ParseFloat(string(line), 64)

				mybdata := IntToBytes(0)
				mybdata = append(mybdata, IntToBytes(2)...)
				mybdata = append(mybdata, Float64ToByte(data)...)
				conn.Write(mybdata)

			} else if mytype == 3 {
				data := string(line)

				mybdata := IntToBytes(0)
				mybdata = append(mybdata, IntToBytes(3)...)
				mybdata = append(mybdata, IntToBytes(len(data))...)
				conn.Write(mybdata)
				conn.Write([]byte(data))

			} else if mytype == 4 {
				data := string(line)
				datalist := strings.Split(data, " # ")
				times, _ := strconv.Atoi(datalist[1])
				password := datalist[0]
				//arr=append(arr,	Pass{datalist[0],times})

				mybdata := IntToBytes(0)
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
		if SmalltoBig {
			mybend := IntToBytes(0)
			mybend = append(mybend, IntToBytes(0)...)
			mybend = append(mybend, IntToBytes(1)...)
			conn.Write(mybend)
		} else {
			mybend := IntToBytes(0)
			mybend = append(mybend, IntToBytes(0)...)
			mybend = append(mybend, IntToBytes(2)...)
			conn.Write(mybend)
		}

	} else {
		mybstart := IntToBytes(1)
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

		//读取文件制定类型发送之
		fi, err := os.Open(sendfilepath)
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
			mytype := datatype
			if mytype == 1 {
				data, _ := strconv.Atoi(string(line))
				//发送整数
				mybdata := IntToBytes(1)
				mybdata = append(mybdata, IntToBytes(1)...)
				mybdata = append(mybdata, IntToBytes(data)...)
				conn.Write(mybdata)

			} else if mytype == 2 {
				data, _ := strconv.ParseFloat(string(line), 64)

				mybdata := IntToBytes(1)
				mybdata = append(mybdata, IntToBytes(2)...)
				mybdata = append(mybdata, Float64ToByte(data)...)
				conn.Write(mybdata)

			} else if mytype == 3 {
				data := string(line)

				mybdata := IntToBytes(1)
				mybdata = append(mybdata, IntToBytes(3)...)
				mybdata = append(mybdata, IntToBytes(len(data))...)
				conn.Write(mybdata)
				conn.Write([]byte(data))

			} else if mytype == 4 {
				data := string(line)
				datalist := strings.Split(data, " # ")
				times, _ := strconv.Atoi(datalist[1])
				password := datalist[0]
				//arr=append(arr,	Pass{datalist[0],times})

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
		if SmalltoBig {
			mybend := IntToBytes(1)
			mybend = append(mybend, IntToBytes(0)...)
			mybend = append(mybend, IntToBytes(1)...)
			conn.Write(mybend)
		} else {
			mybend := IntToBytes(1)
			mybend = append(mybend, IntToBytes(0)...)
			mybend = append(mybend, IntToBytes(2)...)
			conn.Write(mybend)
		}

	}
}

//接收数据
func ReceFile(savefilepath string, conn net.Conn, datatype int) {

	savefile, _ := os.Create(savefilepath)
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
			fmt.Println(data1, data2, data3, "接收到")

			if data2 == 0 && data3 == 0 {
				//开始
				fmt.Println("开始")

			}
			if data2 == 1 {
				//整数
				fmt.Fprintln(save, strconv.Itoa(data3))

			} else if data2 == 2 {
				//实数
				floatnum := ByteToFloat64(buf[16:])
				fmt.Fprintln(save, strconv.FormatFloat(floatnum, 'f', 6, 64)) //arr=append(arr,data3)

			} else if data2 == 3 {
				//字符串

				strbyte := make([]byte, data3, data3)
				length, _ := conn.Read(strbyte)
				if length == data3 { //校验长度
					//arr=append(arr,string(strbyte))
					fmt.Fprintln(save, string(strbyte))
				}

			} else if data2 == 4 {

				buf1 := make([]byte, 8)
				length, _ := conn.Read(buf1)
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

			}

			if data2 == 0 && data3 == 1 {
				//结束，从小到大
				fmt.Println("结束", 1)
				save.Flush()
				savefile.Close()
				return

			}
			if data2 == 0 && data3 == 2 {
				fmt.Println("结束", 2)
				//结束，从大到小
				save.Flush()
				savefile.Close()
				return

			}
		}

	}

}
