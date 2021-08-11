package main

import "fmt"

func main() {
	/*var 变量名 map[keyType]valueType*/
	var mapDemo map[int]string
	mapDemo = make(map[int]string, 10) //key不能重复，value可以重复
	//mapDemo=make(map[int]string)
	mapDemo[1] = "松江"
	mapDemo[2] = "卢俊义"
	fmt.Printf("v=%v\n", mapDemo)
}
