package main

import (
	"fmt"
)

func main() {
	/*var 变量名 map[keyType]valueType*/
	var mapDemo map[int]string
	mapDemo = make(map[int]string, 10) //key不能重复，value可以重复
	//mapDemo=make(map[int]string)
	mapDemo[1] = "松江"  //增
	delete(mapDemo, 1) //删除
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

}
