package main


import (
	"io/ioutil"
	//"github.com/pkg/errors"
	"fmt"
)


type  MyQueue interface {
	Size()int   //大小
	Front()interface{} //第一个元素
	End()interface{} //最后一个元素
	IsEmpty() bool //是否为空
	Enqueue(data interface{}) //入队
	Dequeue() interface{} //出对
	Clear()//清空
}

type  Queue struct{
	datastore []string
	theSize int
}
func(myqueue *Queue) Clear(){
	myqueue.datastore =make([]string,10)  //开辟内存
	myqueue.theSize=0
}

func  NewQueue() *Queue{
	myqueue:=new(Queue)
	myqueue.Clear()
	return myqueue

}
func(myqueue *Queue) Size()int   {
	return  myqueue.theSize //大小
}
func(myqueue *Queue) Front()string {
	if myqueue.Size()==0{ //判断是否为空
		return ""
	}
	return myqueue.datastore[0]
}
func(myqueue *Queue) End()string {
	if myqueue.Size()==0{ //判断是否为空
		return ""
	}
	return myqueue.datastore[myqueue.theSize-1]
}
func(myqueue *Queue) IsEmpty() bool{
	return  myqueue.theSize==0
}
func(myqueue *Queue) Enqueue(data string) {
	//myqueue.datastore=append( myqueue.datastore,data) //入队
	myqueue.datastore[myqueue.theSize]=data
	myqueue.theSize++
}
func(myqueue *Queue) Dequeue() string {
	if myqueue.Size()==0{ //判断是否为空
		return ""
	}
	data:=myqueue.datastore[0]
	if  myqueue.Size()>1 {
		myqueue.datastore = myqueue.datastore[1:myqueue.theSize] //截取
	}
	myqueue.theSize--
	return  data
}

func mainxxxxx(){
	myq:=NewQueue()
	myq.Enqueue("1")
	fmt.Println("s",myq.Dequeue())
	myq.Enqueue("2")
	fmt.Println("s",myq.Dequeue())
	myq.Enqueue("2")
	fmt.Println("s",myq.Dequeue())
}
func  main(){
	files:=[]string{}
	path:="C:\\Users\\Tsinghua-yincheng\\Desktop\\day4"





	myq:=NewQueue()
	myq.Enqueue(path)

	for ;;{
		path =myq.Dequeue()
		if path==""{
			break
		}
		//files=append(files,path)
		fmt.Println("get",path)
		read,err:=ioutil.ReadDir(path)//读取文件夹
		if err!=nil{
			break
		}
		for _,fi:=range read{
			if fi.IsDir(){
				fulldir:=path+"\\"+fi.Name()//文件夹路径
				fmt.Println("setdir",fulldir)
				//myq.Enqueue(fulldir)
				myq.Enqueue(fulldir)
			}else{
				fullname:=path+"\\"+fi.Name()//处理文件
				fmt.Println("setfile",fullname)
			}
		}

	}
	for i:=0;i<len(files);i++{
		fmt.Println(files[i])
	}



}

