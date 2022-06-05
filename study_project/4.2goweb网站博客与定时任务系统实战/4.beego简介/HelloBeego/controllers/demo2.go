package controllers

import (
	"github.com/astaxie/beego"
)

type Demo2Controller struct {
	beego.Controller
}

//http://localhost:8080/demo2/aa?username=xdl&password=123
func (this *Demo2Controller) GetFunc() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Ctx.WriteString("username = " + username + "\npassword = " + password + "\n用于处理get请求")
	//fmt.Println("username = " + username + "password = " + password)
}

//http://localhost:8080/demo2/aa?username=xdl&password=123
func (this *Demo2Controller) PostFunc() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Ctx.WriteString("username = " + username + "\npassword = " + password + "\n用于处理post请求")
	//fmt.Println("username = " + username + "password = " + password)
}

//http://localhost:8080/demo2/aa?username=xdl&password=123
func (this *Demo2Controller) GetAndPostFunc() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Ctx.WriteString("username = " + username + "\npassword = " + password + "\n用于处理post和get请求")
	//fmt.Println("username = " + username + "password = " + password)
}

//http://localhost:8080/demo2/aa?username=xdl&password=123
func (this *Demo2Controller) AnyFunc() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Ctx.WriteString("username = " + username + "\npassword = " + password + "\n用于处理任意请求")
	//fmt.Println("username = " + username + "password = " + password)
}
