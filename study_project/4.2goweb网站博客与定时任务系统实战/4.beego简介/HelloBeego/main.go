package main

import (
	_ "HelloBeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.ViewsPath = "myview"
	beego.Run()
}
