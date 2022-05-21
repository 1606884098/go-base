package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 84331445

func main() {
	allstrs := make([]string, N, N) //初始化数组

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1数据结构\\QQ.txt"
	QQfile, _ := os.Open(path)    //打开文件
	defer QQfile.Close()          //最后关闭文件
	i := 0                        //统计一共多少行
	br := bufio.NewReader(QQfile) //读取数据
	for {
		line, _, end := br.ReadLine() //读取一行数据
		if end == io.EOF {            //文件关闭。跳出循环
			break
		}
		allstrs[i] = string(line)
		i++
	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 20)

	for {
		fmt.Println("请输入要查询的数据")
		var inputstr string = "yincheng"
		fmt.Scanln(&inputstr)

		starttime := time.Now() //时间开始
		for j := 0; j < N; j++ {

			if strings.Contains(allstrs[j], inputstr) { //字符串搜索
				fmt.Println(allstrs[j])
			}
		}
		fmt.Println("一共用了", time.Since(starttime))
		//break
	}

}

//一共用了 2.7781436s
/*
81829191----kekexili2008
33045139----woaikekexili
174537164----kekexili0726
243383050----kekexili
297309366----52kekexili
290666632----kekexili520
249586779----kekexili828Bi
289728890----kekexili123
641238655----kekexilimanyi00
632966515----kekexiliqhhhzh
516607963----woaikekexili
137412365----kekexili493+3721
504914917----kekexili886
342441680----kekexili132412
823907904----kekexili321118
371986625----kekexili317
754802260----kekexili
734293663----kekexili
584716165----kekexili
756297451----kekexili5213213
546801633----kekexili090214
1959184220----123456kekexili
645070807----kekexili3307


*/
