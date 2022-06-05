package main


import (
	"net"
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
	"strconv"
	"./pipelineMiddleWare"
)





func IntTobytes(n int)[]byte{
	data:=int64(n)
	bytebuffer:=bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer,binary.BigEndian,data)
	return bytebuffer.Bytes()
}
func  BytesToInt(bts []byte)int{
	bytebuffer:=bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer,binary.BigEndian,&data)
	return int(data)
}

func ServerMsgHandler(conn net.Conn) <-chan int {
	out:=make(chan int,1024)//新的管道

	buf :=make([]byte,16)
	arr:= []int{} //数组保存数据
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("Sever  close")
			return nil
		}
		if n==16 {
			data1:=BytesToInt(buf[:len(buf)/2])//取出的第一个数据
			data2:=BytesToInt(buf[len(buf)/2:])//取出第二个数据
			if data1==0 && data2==0{
				arr=make([]int,0,0) //开辟数据
			}
			if data1==1 {
				arr=append(arr,data2)
			}
			if data1==0 && data2==1{
				fmt.Println("数组接收完成",arr)
				for i:=0;i<len(arr);i++{
					out <-arr[i]  //数组压入管道
				}
				close(out) //关闭管道
				return out

				arr=make([]int,0,0) //开辟数据
			}



		}




	}

	return nil

}



func SendArray(arr []int ,conn net.Conn){
	length:=len(arr)
	mybstart:=IntTobytes(0)
	mybstart=append(mybstart,IntTobytes(0)...)
	conn.Write(mybstart)

	for i:=0;i<length;i++{
		mybdata:=IntTobytes(1)
		mybdata=append(mybdata,IntTobytes(arr[i])...)
		conn.Write(mybdata)
	}

	mybend:=IntTobytes(0)
	mybend=append(mybend,IntTobytes(1)...)
	conn.Write(mybend)
}



func main(){
	arrlist:=[][]int{{1,9,2,8,7,3,5,6,10,4,23,24},{11,19,12,18,17,13,15,16,101,14,123,124}}
	sortResults:=[]<-chan int{}



	for i:=0;i<2;i++{
		tcpaddr,err:=net.ResolveTCPAddr("tcp4","127.0.0.1:700"+strconv.Itoa(1+i))
		if err!=nil{
			panic(err)
		}
		conn,err:=net.DialTCP("tcp",nil,tcpaddr)//链接
		if err!=nil{
			panic(err)
		}


		SendArray(arrlist[i],conn)
		sortResults=append(sortResults,ServerMsgHandler(conn))

	}
	fmt.Println(len(sortResults))
	last:=pipelineMiddleWare.Merge(sortResults[0],sortResults[1])
	for v:=range last{
		fmt.Printf("%d ",v)
	}
	time.Sleep(time.Second*30)



}

