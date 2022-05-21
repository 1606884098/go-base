package main

import "./Chain"
import "./Observer"
import (
	"./Interpreter"
	"fmt"
)
import "./State"

func main() {
	ctx := State.NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}
	for i := 0; i < 98; i++ {
		todayAndNext()
	}

}
func main3() {
	//1+2
	//1+2+3
	//1+2+3+5-4
	p := &Interpreter.Parser{}
	fmt.Print("start\n")
	p.Parse("1 + 3 - 2 + 5 - 6 + 7 - 10")
	fmt.Println(p.Result().Interpret())

}

func main2() {
	subject := Observer.NewSubject()
	r1 := Observer.NewReader("lixiang")
	r2 := Observer.NewReader("hupeng")
	r3 := Observer.NewReader("zhangbo")
	subject.Attch(r1)
	subject.Attch(r2)
	subject.Attch(r3)
	subject.UpdateContext("妹子来了")
	r4 := Observer.NewReader("yangjie")
	subject.Attch(r4)
	subject.UpdateContext("漂亮妹子来了")
}

func main1() {
	c1 := Chain.NewProjectManagerChain()
	c2 := Chain.NewDepManagerChain()
	c3 := Chain.NewGenaralManagerChain()

	c1.SetSuceessor(c2)
	c2.SetSuceessor(c3)

	var c Chain.Manger = c1
	c.HandleFeeRequest("hupeng", 1500)
	c.HandleFeeRequest("hupeng", 500)
	c.HandleFeeRequest("lixiang", 4500)
	c.HandleFeeRequest("lixiang", 11500)
	c.HandleFeeRequest("weishangyin", 4500)
	c.HandleFeeRequest("weishangyin", 11500)
}
