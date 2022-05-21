package main

import "fmt"

//适配的目标接口
type Target interface {
	Request(int, int) string
}
type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

//接口包装
func (adap *adapter) Request(a, b int) string {
	return adap.SpecficRequest(a, b)
}

//被适配
type Adaptee interface {
	SpecficRequest(int, int) string
}

//载体，被适配的目标类
type adapeeImpl struct{}

//工厂函数
func NewAdaptee() Adaptee {
	return &adapeeImpl{}
}

//实际函数
func (ada *adapeeImpl) SpecficRequest(a, b int) string {
	fmt.Println(a, b)
	return "SpecficRequest你妹"
}

func main() {
	adapee := NewAdaptee()       //适配器
	target := NewAdapter(adapee) //传递进入
	res := target.Request(1, 2)
	fmt.Println(res)

}
