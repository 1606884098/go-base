package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name    string `json:"id"`
	Age     int    `json:"age"`               //通过指定tag实现json序列化该字段时的key
	Address string `json:"address,omitempty"` //omitempty过滤空值
}

/*func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}*/
func (p Person) testMathod() { //结构体绑定方法
	p.name = "值类型副本被修改"
	fmt.Println(p.name)
}

func (p *Person) testmathod1() {
	p.name = "指针类型接收者被修改"
	fmt.Println("v=", p.name)
}

/*结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）*/
func NewPerson(name string) *Person { //构造方法
	return &Person{name: name}
}

func main() {
	type myint = int64 //myint相当于是别名
	var c myint = 10
	fmt.Println(c)

	var p Person
	p.name = "jack"
	p.testMathod() //调用方法
	p.testmathod1()

	p1 := Person{"张三", 10, "中国"}
	fmt.Printf("p1=%T\n", p1)

	//指针类型的结构体
	var p3 = new(Person)
	p3.name = "hello"
	fmt.Println(p3)        //&{hello}
	fmt.Printf("%T\n", p3) //*main.Person

	p4 := &Person{}        //&Person{name:"dd"}
	fmt.Println(p4)        //&{}
	fmt.Printf("%T\n", p4) //*main.Person
	p4.name = "dd"
	fmt.Println(p4) //&{dd}

	var user struct { //匿名结构体
		name string
		age  int
	}
	user.name = "老王"
	user.age = 10
	fmt.Println(user)
	//实现结构体的构造函数
	structFunc := NewPerson("实现构造函数")
	fmt.Printf("structFunc=%T\n", structFunc)

	/*	指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，
		在方法结束后，修改都是有效的。*/
	p5 := new(Person)
	p5.name = "指针接收者"
	p5.testmathod1()
	fmt.Println(p5.name)
	/*	当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获
		取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。*/
	p6 := new(Person)
	p6.name = "值类型本身"
	p6.testMathod()      //值类型副本被修改
	fmt.Println(p6.name) //值类型本身

	//匿名字段的结构体
	niming()

	//值类型相同就是同一个对象 如：
	p7 := Person{"11", 1, "aa"}
	p8 := Person{"11", 1, "aa"}
	fmt.Println(p7 == p8) //true
	//指针类型指向不同地址
	p9 := &Person{"11", 1, "aa"}
	p10 := &Person{"11", 1, "aa"}
	fmt.Println(p9 == p10) //flase

}

type mycc int
type Person1 struct { //匿名字段的结构体，可以理解是不同类型的数据组合，类型可以是自定义
	Person //实现了继承
	string
	int
	mycc
}

func niming() {
	p1 := new(Person1)
	p1.string = "niming"
	p1.int = 10
	p1.mycc = 10
	p1.name = "继承后的名字"
	p1.testMathod()
	fmt.Println(p1)

	//结构体序列化成json
	p1.Age = 199
	data, err := json.Marshal(p1) //定义的字段一定要是大写才能转化json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("json:%s\n", data)

	//json反序列化成结构体
	str := `{"Age":999}`
	srtPerson := new(Person)
	err1 := json.Unmarshal([]byte(str), srtPerson)
	if err1 != nil {
		fmt.Println("解码失败！", err1)
		return
	}
	fmt.Println(srtPerson.Age)

}
