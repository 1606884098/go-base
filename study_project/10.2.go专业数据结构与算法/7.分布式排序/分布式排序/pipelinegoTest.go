package main

import "os"
import (
	"./pipelineMiddleWare"
	"bufio"
	"fmt"
	"time"
)

func main(){
	var filename="data1.in"//文件写入
	var count=100000
	file ,err:=os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()//延迟关闭文件

	mypipe:=pipelineMiddleWare.RandomSource(count) //管道装随机数
	writer:=bufio.NewWriter(file)//写入
	pipelineMiddleWare.WriterSlink(writer,mypipe)//写入
	writer.Flush()//刷新


	file ,err=os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()//延迟关闭文件
	mypipeX:=pipelineMiddleWare.ReaderSource(bufio.NewReader(file),-1)
	counter:=0
	for v:=range 	mypipeX{
		fmt.Println(v)
		counter++
		if counter>1000{
			break
		}
	}



}

func main2x(){
	go func() {
		myp:=pipelineMiddleWare.Merge(
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(3,9,2,1,10)),
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(13,19,12,11,110)),
		)
		for v:=range myp{
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second*10)
}
