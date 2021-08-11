package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

/*
	sha1 sha256 sha512 md5都是常用的hash函数（算法）
*/
func main() {
	var a = "长江后浪推前浪"
	md5run := md5.New()
	md5run.Write([]byte(a))
	result := md5run.Sum(nil)
	fmt.Printf("%x\n", result)

	/*
		sha512出现hash碰撞的概率为十万亿分之一，百度就用这个方法来判断这一个文件是不是同一个文件，
		来实现秒传
		也可以用来判断一个文件有没被修改过里面的内容
	*/
	md5run1 := sha512.New()
	md5run1.Write([]byte(a))
	result1 := md5run1.Sum(nil)
	fmt.Printf("%x\n", result1)

	path := `E:\wltool\tool\笔记.doc` //如果这个文件修改了一个字节sha512的结果都不一样的
	file, _ := os.Open(path)        //打开文件
	md5run2 := sha512.New()
	io.Copy(md5run2, file)               //拷贝数据
	fmt.Printf("%x\n", md5run2.Sum(nil)) //计算hash值

}
