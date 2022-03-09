package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

//00
//12
//13
//19
//01  //整数类型   字节
//实数类型  字节
//0 0 0  //字符串   字节
//1 4  calc
//1 5 cmdtr
//0 0 1

func IntToBytes(n int) []byte {
	data := int64(n)                                 //数据类型转换
	bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
	binary.Write(bytebuffer, binary.BigEndian, data) //按照二进制写入字节
	return bytebuffer.Bytes()                        //返回字节结合
}

func BytesToInt(bs []byte) int {
	bytebuffer := bytes.NewBuffer(bs) //根据二进制写入二进制结合
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data) //解码
	return int(data)
}

func ByteToFloat32(bs []byte) float32 {
	bits := binary.LittleEndian.Uint32(bs) //解码
	return math.Float32frombits(bits)
}
func ByteToFloat64(bs []byte) float64 {
	bits := binary.LittleEndian.Uint64(bs) //解码
	return math.Float64frombits(bits)
}

//浮点数转化为字节
func Float32ToByte(data float32) []byte {
	bits := math.Float32bits(data) //math的方法
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits) //填充
	return bytes
}
func Float64ToByte(data float64) []byte {
	bits := math.Float64bits(data) //math的方法
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits) //填充
	return bytes
}

func main() {
	//字符串转字节
	mystr := "calc"
	fmt.Println([]byte(mystr))
	fmt.Println(string([]byte(mystr)))
	//整数类型   字节
	//实数类型  字节
	fmt.Println(IntToBytes(1233231))
	fmt.Println(BytesToInt(IntToBytes(1233231)))
	//整数类型   字节
	//实数类型  字节
	fmt.Println(Float32ToByte(11221.2346567))
	fmt.Println(ByteToFloat32(Float32ToByte(11221.2346567)))
	fmt.Println(Float64ToByte(11221.2346567))
	fmt.Println(ByteToFloat64(Float64ToByte(11221.2346567)))
}
