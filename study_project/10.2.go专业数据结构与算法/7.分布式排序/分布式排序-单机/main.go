package main
//1.本地归并排序，2.多线程 3.分布式

//生成随机数组
import (
	"./pipelineMiddleWare"
	"os"
	"bufio"
	"fmt"
	"time"
)


//多线程-调用中间件完成
func createPipeline(filename string ,filesize int ,chunkCount int)<-chan int{


	//var filename="data1.in"//文件写入
	//var count=100000
	file ,err:=os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()//延迟关闭文件

	mypipe:=pipelineMiddleWare.RandomSource(filesize/8) //管道装随机数
	writer:=bufio.NewWriter(file)//写入
	pipelineMiddleWare.WriterSlink(writer,mypipe)//写入
	writer.Flush()//刷新


	chunkSize:=filesize/chunkCount //数量
	sortResults:=[]<-chan int{} //排序结果，一个数组，每一个元素是个管道
	pipelineMiddleWare.Init()//初始化
	file,err=os.Open(filename)//打开文件
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	for  i:=0;i<chunkCount;i++{

		file.Seek(int64(i*chunkSize),0)//跳到文件指针
		source:=pipelineMiddleWare.ReaderSource(bufio.NewReader(file),chunkSize)//读取
		sortResults=append(sortResults,pipelineMiddleWare.InMemorySort(source))//结果排序


	}
	return  pipelineMiddleWare.MergeN(sortResults...)


}
//写入文件
func writetofile(in <-chan int,filename string){
	file ,err:=os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	writer:=bufio.NewWriter(file)
	defer writer.Flush()//刷新

	pipelineMiddleWare.WriterSlink(writer,in) //写入数据


}
//显示文件
func showfile(filename string){
	file,err:=os.Open(filename)//打开文件
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	p:=pipelineMiddleWare.ReaderSource(bufio.NewReader(file),-1)


	counter:=0
	for v:=range 	p{
		fmt.Println(v)
		counter++
		if counter>1000{
			break
		}
	}


}

func main(){
	go func() {
		time.Sleep(time.Second*1000)
	}()
	p:=createPipeline("C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\big.in",800000,4)
	writetofile(p,"C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\big.out")
	showfile("C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\big.out")



}

