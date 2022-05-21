package main

import "fmt"

//原型对象需要实现的接口
//map  map[""1232]=1232
//拷贝原有的对象

type Cloneable interface {
	Clone() Cloneable
}

//原型对象的类
type PrototypeManger struct {
	prototypes map[string]Cloneable
}

//构造初始化
func NewPrototypeManger() *PrototypeManger {
	return &PrototypeManger{make(map[string]Cloneable)}
}

//抓取
func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}

//设置
func (p *PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	//tc:=*t
	//return &tc 深复制
	return t
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t   //开辟内存新建变量，存储指针指向的内容
	return &tc //返回地址
}

func main() {
	mgr := NewPrototypeManger()
	t1 := &Type1{name: "type1"}
	mgr.Set("t1", t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone() //复制
	if t11 == t22 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}

}
