package controllers

import "github.com/astaxie/beego"

type Demo4Controller struct {
	beego.Controller
}

func (this *Demo4Controller) URLMpping() {
	this.Mapping("GetFunc", this.GetFunc)
	this.Mapping("PostFunc", this.PostFunc)
}

//@router /demo4/getfunc/:key [get]
func (this *Demo4Controller) GetFunc() {
	key := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString("key = " + key + "，用于处理get请求!")
}

//http://localhost:8080/demo4/getfunc/123

//@router /demo4/postfunc/:key [post]
func (this *Demo4Controller) PostFunc() {
	key := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString("key = " + key + "，用于处理get请求!")
}
