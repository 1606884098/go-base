package controllers

import "github.com/astaxie/beego"

type CookieController struct {
	beego.Controller
}

func (this *CookieController) Get() {
	if this.Ctx.GetCookie("username") == "" {
		this.Ctx.SetCookie("username", "admin", nil, "/cookie")
		this.Ctx.WriteString("cookie设置成功!")
	} else {
		username := this.Ctx.GetCookie("username")
		this.Ctx.WriteString("username = " + username)
	}
}
