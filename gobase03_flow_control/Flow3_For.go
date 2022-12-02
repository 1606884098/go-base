package main

import "fmt"

/*
我们有专门实现这种循环的结构就是for结构（GO语言中只有for循环结构，没有while,do-while
结构），基本语法结构如下：

for 表达式1;表达式2;表达式3{
	循环体
}
表达式1:定义一个循环的变量，记录循环的次数
表达式2：一般为循环条件，循环多少次
表达式3：一般为改变循环条件的代码，使循环条件终有一天不再成立
循环体：重复要做的事情。


for 循环条件 condition {
    // 循环体代码
}

for {
    // 循环体代码
}

for 循环的 range 格式对 string、slice、array、map、channel 等进行迭代循环。
array、slice、string 返回索引和值；map 返回键和值；channel 只返回通道内的值。
其语法结构如下所示。

for key, value := range oldMap {
    newMap[key] = value
}

for 嵌套循环语句
Go语言允许在循环体内使用循环。其语法结构如下所示。
for [condition | (init; condition; increment) | Range] {
    for [condition | (init; condition; increment) | Range] {
        statement(s);
    }
    statement(s);
}


关键字：
break 语句用来跳出循环体，终止当前正在执行的 for 循环，并开始执行循环之后的语句
continue  跳过当前循环，执行下一次循环语句
return 也可以结束一个循环
*/
func main() {
	var str = "woai北京"
	forFunc(str)
	forRangeFunc(str)
}

/*对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节*/
func forFunc(s string) {
	fmt.Println("普通for----------")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i]) //出现乱码
	}
}

/*for-range 遍历方式而言，是按照字符方式遍历。因此如果有字符串有中文，也是 ok*/
func forRangeFunc(s string) {
	fmt.Println("for-range----------")
	for index, val := range s {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}
