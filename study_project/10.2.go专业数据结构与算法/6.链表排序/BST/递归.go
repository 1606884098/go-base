package main

import (
"fmt"
"github.com/pkg/errors"
)

type  Node struct{
	data  int
	next * Node
}

type  LinkStack interface {
	IsEmpty() bool
	Push(value  int)
	Pop()( int,error)
	Length() int
}

func NewStack()*Node{
	return &Node{}
}

func (n*Node)IsEmpty() bool{  //判断是否为空
	return n.next==nil
}
func (n*Node)Push(value  int){
	newnode:=&Node{data:value} //初始化
	newnode.next=n.next
	n.next=newnode
}
func (n*Node)Pop()( int,error){
	if n.IsEmpty()==true{
		return -1,errors.New("bug")
	}
	value :=n.next.data
	n.next=n.next.next
	return  value,nil
}
func (n*Node)Length() int{
	pnext:=n
	length:=0
	for pnext.next!=nil{//返回长度
		pnext=pnext.next
		length++
	}
	return length
}

func Add(num int ) int {
	if num==0{
		return 0
	}else{
		return  num+Add(num-1)
	}
}
func FAB(num int)int {
	if num==1 ||num==2{
		return 1
	}else{
		return FAB(num-1)+FAB(num-2)
	}
}
//5  f(4) +f(3)   f(2)+f(1)+f(2)+f(2)+f(1))
//1  1  2   3   5  8  13  21  34  55
func main(){
	//fmt.Println(Add(10))
	mystack:=NewStack()
	mystack.Push(10)
	last:=0
	for !mystack.IsEmpty(){
		data,err:=mystack.Pop()
		if err!=nil{
			break
		}
		if data==1 || data==2{
			last+=1
		}else{
			mystack.Push((data-1))
			mystack.Push((data-2))
		}
	}

	fmt.Println(last)

}



func  mainx(){
	mystack:=NewStack()
	for i:=0;i<10;i++{
		mystack.Push(i)
	}
	for data,err:=mystack.Pop();err==nil;data,err=mystack.Pop(){
		fmt.Println(data)
	}




}
