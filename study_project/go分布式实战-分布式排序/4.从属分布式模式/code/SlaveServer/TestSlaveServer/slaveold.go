package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

func DiskfileSortAndSend(filepath string, mytype int, isSmalltoBig bool, conn net.Conn) {

	if mytype == 1 {
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
			data, _ := strconv.Atoi(string(line)) //
			arr = append(arr, data)
		}
		//排序
		mydata := new(QuickSortData)
		mydata.IsSmalltoBig = true
		mydata.Data = arr
		QuickSortDataByType(mydata, mytype)
		mydata.QuickSort()
		fmt.Println("最终数组", arr)

		//写入
		newpath := strings.Replace(filepath, ".txt", "sort.txt", -1)
		savefile, _ := os.Create(newpath)
		save := bufio.NewWriter(savefile) //保存
		for i := 0; i < len(mydata.Data); i++ {
			fmt.Fprintln(save, strconv.Itoa(mydata.Data[i].(int)))
		}
		save.Flush()
		savefile.Close()
		//释放内存
		arr = nil
		runtime.GC()
		debug.FreeOSMemory() //释放内存

		//发送
		//整数
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

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
			data, _ := strconv.Atoi(string(line)) //
			//arr=append(arr,	data)
			conn.Write(IntToBytes(0))
			mybdata := IntToBytes(1)
			conn.Write(mybdata)
			conn.Write(IntToBytes(data))
		}
		fi.Close() //打开文件
		/*
			for i:=0;i<len(mydata.Data);i++{

				conn.Write(IntToBytes(0))
				mybdata:=IntToBytes(1)
				conn.Write(mybdata)
				conn.Write(IntToBytes(mydata.Data[i].(int)))
			}*/

		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
		}
		conn.Write(mybend)

	} else if mytype == 2 {
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
			data, _ := strconv.ParseFloat(string(line), 64)
			arr = append(arr, data)
		}
		//排序
		mydata := new(QuickSortData)
		mydata.IsSmalltoBig = true
		mydata.Data = arr
		QuickSortDataByType(mydata, mytype)
		mydata.QuickSort()
		fmt.Println("最终数组", arr)

		//写入
		newpath := strings.Replace(filepath, ".txt", "sort.txt", -1)
		savefile, _ := os.Create(newpath)
		save := bufio.NewWriter(savefile) //保存
		for i := 0; i < len(mydata.Data); i++ {
			fmt.Fprintln(save, strconv.FormatFloat(mydata.Data[i].(float64), 'f', 6, 64))
		}
		save.Flush()
		savefile.Close()
		//释放内存
		arr = nil
		runtime.GC()
		debug.FreeOSMemory() //释放内存

		//发送
		//整数
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

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
			data, _ := strconv.ParseFloat(string(line), 64)
			//arr=append(arr,	data)
			conn.Write(IntToBytes(0))
			mybdata := IntToBytes(2)
			conn.Write(mybdata)
			conn.Write(Float64ToByte(data))
		}
		fi.Close() //打开文件
		/*
			for i:=0;i<len(mydata.Data);i++{

				conn.Write(IntToBytes(0))
				mybdata:=IntToBytes(1)
				conn.Write(mybdata)
				conn.Write(IntToBytes(mydata.Data[i].(int)))
			}*/

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
			data := string(line)
			arr = append(arr, data)
		}
		//排序
		mydata := new(QuickSortData)
		mydata.IsSmalltoBig = true
		mydata.Data = arr
		QuickSortDataByType(mydata, mytype)
		mydata.QuickSort()
		fmt.Println("最终数组", arr)

		//写入
		newpath := strings.Replace(filepath, ".txt", "sort.txt", -1)
		savefile, _ := os.Create(newpath)
		save := bufio.NewWriter(savefile) //保存
		for i := 0; i < len(mydata.Data); i++ {
			fmt.Fprintln(save, mydata.Data[i].(string))
		}
		save.Flush()
		savefile.Close()
		//释放内存
		arr = nil
		runtime.GC()
		debug.FreeOSMemory() //释放内存

		//发送
		//整数
		mybstart := IntToBytes(0) //内存
		mybstart = append(mybstart, IntToBytes(0)...)
		mybstart = append(mybstart, IntToBytes(0)...)
		conn.Write(mybstart)

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

			//arr=append(arr,	data)
			conn.Write(IntToBytes(0))
			mybdata := IntToBytes(3)
			conn.Write(mybdata)
			mybdata1 := IntToBytes(len(string(line)))
			conn.Write(mybdata1)
			conn.Write(line)
		}
		fi.Close() //打开文件
		/*
			for i:=0;i<len(mydata.Data);i++{

				conn.Write(IntToBytes(0))
				mybdata:=IntToBytes(1)
				conn.Write(mybdata)
				conn.Write(IntToBytes(mydata.Data[i].(int)))
			}*/

		//结束
		mybend := IntToBytes(0)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
		}
		conn.Write(mybend)

	} else if mytype == 4 {

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
			data := string(line)
			datalist := strings.Split(data, " # ")
			times, _ := strconv.Atoi(datalist[1])
			arr = append(arr, Pass{datalist[0], times})
		}
		//排序
		mydata := new(QuickSortData)
		mydata.IsSmalltoBig = true
		mydata.Data = arr
		QuickSortDataByType(mydata, mytype)
		mydata.QuickSort()
		fmt.Println("最终数组", arr)

		//写入
		newpath := strings.Replace(filepath, ".txt", "sort.txt", -1)
		savefile, _ := os.Create(newpath)
		save := bufio.NewWriter(savefile) //保存
		for i := 0; i < len(mydata.Data); i++ {
			fmt.Fprintln(save, mydata.Data[i].(Pass).PassWord+" # "+strconv.Itoa(mydata.Data[i].(Pass).times))
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
			data := string(line)
			datalist := strings.Split(data, " # ")
			password := datalist[0]
			times, _ := strconv.Atoi(datalist[1])

			//arr=append(arr,	data)
			conn.Write(IntToBytes(1))
			mybdata := IntToBytes(4)
			conn.Write(mybdata)
			conn.Write(IntToBytes(times))
			conn.Write(IntToBytes(len(password)))
			//104
			conn.Write([]byte(password))

			mybdata1 := IntToBytes(len(string(line)))
			conn.Write(mybdata1)
			conn.Write(line)

		}
		fi.Close() //打开文件
		/*
			for i:=0;i<len(mydata.Data);i++{

				conn.Write(IntToBytes(0))
				mybdata:=IntToBytes(1)
				conn.Write(mybdata)
				conn.Write(IntToBytes(mydata.Data[i].(int)))
			}*/

		//结束
		mybend := IntToBytes(1)
		mybend = append(mybend, IntToBytes(0)...)
		if isSmalltoBig {
			mybend = append(mybend, IntToBytes(1)...)
		} else {
			mybend = append(mybend, IntToBytes(2)...)
		}
		conn.Write(mybend)
	}
}
