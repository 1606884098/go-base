package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Demo3Controller struct {
	beego.Controller
}

func (this *Demo3Controller) GetFunc() {
	params := this.Ctx.Input.Params()
	for key, value := range params {
		fmt.Println(key, " = ", value)
	}

	fmt.Println("params[splat] = ", params[":splat"])

	this.Ctx.WriteString("我用于处理get请求!")
}

func (this *Demo3Controller) PostFunc() {
	params := this.Ctx.Input.Params()
	for key, value := range params {
		fmt.Println(key, " = ", value)
	}

	fmt.Println("params[splat] = ", params[":splat"])

	this.Ctx.WriteString("我用于处理post请求!")
}
