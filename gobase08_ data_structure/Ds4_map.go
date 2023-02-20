package main

import (
	"fmt"
	_ "hash/maphash"
)

/*
type hmap struct{//src/runtime/map.go
在一个map里所有的键都是唯一的，而且必须是支持==和!=操作符的类型，切片、函数以及包含切片的结构类型这些类型由于
具有引用语义，不能作为映射的键，使用这些类型会造成编译错误：
map值可以是任意类型，没有限制。map里所有键的数据类型必须是相同的，值也必须如此，但键和值的数据类型可以不相同。*/
func main() {
	/*var 变量名 map[keyType]valueType*/
	var mapDemo map[int]string
	mapDemo = make(map[int]string, 10) //key不能重复，value可以重复
	//mapDemo=make(map[int]string)
	mapDemo[1] = "松江"     //增
	mapDemo[11] = "value" //如果 key 还没有，就是增加，如果 key 存在就是修改。
	delete(mapDemo, 1)    //删除
	mapDemo[2] = "卢俊义"
	mapDemo[2] = "潘小姐" //改
	fmt.Printf("v=%v\n", mapDemo)
	var m1 map[int]string = map[int]string{1: "Luffy", 2: "Sanji"}
	m2 := map[int]string{1: "Luffy", 2: "Sanji"}
	valueD := m1[1] //直接查
	fmt.Printf("key 1的值=%v\n", valueD)
	value, ok := m1[1] //查
	fmt.Printf("key 1的值=%v,是否存在=%v\n", value, ok)
	fmt.Printf("m1的值=%v,m2的值=%v\n", m1, m2)

	for k, v := range m1 {
		fmt.Printf("%d ----> %s\n", k, v)
		//1 ----> Luffy
		//2 ----> yoyo
	}

	for k := range m1 {
		fmt.Printf("%d ----> %s\n", k, m1[k])
		//1 ----> Luffy
		//2 ----> Sanji
	}
	//map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value6, ok := sliceMap[key]
	if !ok {
		value6 = make([]string, 0, 2)
	}
	value6 = append(value6, "北京", "上海")
	sliceMap[key] = value6
	fmt.Println(sliceMap)

}
