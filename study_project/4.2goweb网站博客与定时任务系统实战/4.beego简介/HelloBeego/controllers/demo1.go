package controllers

import "github.com/astaxie/beego"

type Demo1Controller struct {
	beego.Controller
}

func (this *Demo1Controller) Get() {
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("id = " + id)

	/*	username := this.Ctx.Input.Param(":username")
		this.Ctx.WriteString("username = " + username)*/

	/*	ext := this.Ctx.Input.Param(":ext")//后缀名
		path := this.Ctx.Input.Param(":path")//文件名
		this.Ctx.WriteString("ext = " + ext + "\npath = " + path)*/
}
