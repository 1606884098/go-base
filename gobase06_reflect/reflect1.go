package main
/*
反射是指在程序运行期间对程序本身进行访问和修改的能力。
程序在编译时，变量被转换为内存地址，变量名不会被编译器
写入到可执行部分。在运行程序时，程序无法获取自身的信息。

Go语言中的变量是分为两部分的:
类型信息：预先定义好的元信息。
值信息：程序运行过程中可动态变化的

reflect包
在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的
reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
 */
import (
	"fmt"
	"reflect"
)

type myInt int64
/*
type name和type kind
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用
type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分
指针、结构体等大品种的类型时，就会用到种类（Kind）。 举个例子，我们定义了两个指针类型
和两个结构体类型，通过反射查看它们的类型和种类。
Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。

type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
 */
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}
/*
reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
reflect.Value与原始值之间可以互相转换。
reflect.Value类型提供的获取原始值的方法如下：
Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
Bool() bool	将值以 bool 类型返回
Bytes() []bytes	将值以字节数组 []bytes 类型返回
String() string	将值以字符串类型返回
 */
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
/*
想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修
改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
 */
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct


	var f float32 = 3.14
	var g int64 = 100
	reflectValue(f) // type is float32, value is 3.140000
	reflectValue(g) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	h := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", h) // type c :reflect.Value

	var i int64 = 100
	// reflectSetValue1(i) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&i)
	fmt.Println(i)

	/*
	isNil()和isValid():
	IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。

	func (v Value) IsNil() bool
	IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、
	切片之一；否则IsNil函数会导致panic
	unc (v Value) IsValid() bool
	IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、
	Kind之外的方法都会导致panic。
	 */
	// *int类型空指针
	var j *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(j).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	k := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(k).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(k).MethodByName("abc").IsValid())
	// map
	l := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(l).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

}