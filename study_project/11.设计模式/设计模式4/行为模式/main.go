package main

import "./Chain"
import "./Observer"
import (
	"./Interpreter"
	"fmt"
)
import "./State"
import "./Template"
import "./Mediator"
import "./Strategy"
import "./MEMENTO"
import "./Command"
import "./VISITOR"

func main() {
	c := VISITOR.CustomerCol{}
	c.Add(VISITOR.NewEnterpriseCustomer("Microsoft"))

	c.Add(VISITOR.NewIndividualCustomer("billgates"))
	c.Add(VISITOR.NewEnterpriseCustomer("Google"))
	//c.Accept(&VISITOR.SerivceRequesttVisitor{} )
	c.Accept(&VISITOR.AnalysisVisitor{})
}
func main9() {
	mb := &Command.MotherBoard{}
	cmd1 := Command.NewMMCommand1(mb)
	cmd2 := Command.NewMMCommand2(mb)

	box1 := Command.NewBox(cmd1, cmd2)
	box1.GoWarmBed()
	box1.GoWashclothes()

	//box2:=Command.NewBox(cmd1,cmd2)
	//box2.GoWarmBed()
	//box2.GoWashclothes()

}
func main8() {
	var xiaoxianglong *MEMENTO.MMfeel = &MEMENTO.MMfeel{0, 0, 0}
	xiaoxianglong.A第一次见面(170, 500000, 70)
	xiaoxianglong.A中彩票()
	xiaoxianglong.A去韩国()
	xiaoxianglong.A断腿骨()
	xiaoxianglong.A妹子的分数()
	fmt.Println(MEMENTO.States)

}
func main7() {
	//ctx:=Strategy.NewMMContext("marry",18,&Strategy.Girl{})
	ctx := Strategy.NewMMContext("alis", 28, &Strategy.Women{})
	//if age<18  facotry

	ctx.Pao()
}
func main6() {
	meditor := Mediator.GetMediatorInstace()
	fmt.Println(meditor)
	meditor.Ccpu.Process("hello")
	meditor.Cgpu.Display("hello")
	meditor.Cdisk.Store("hello")
	meditor.Cmem.Dump("gogogo")
	meditor.Changed(meditor.Ccpu)
	meditor.Changed(meditor.Cmem)

}

func main5() {
	var downloader Template.Downloader = Template.NewHttpDownLoader()
	downloader.Download("http://tsinghua.edu.cn/ai.zip")
	var downloader1 Template.Downloader = Template.NewFtpDownLoader()
	downloader1.Download("ftp://tsinghua.edu.cn/ai.zip")
}

func main4() {
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
