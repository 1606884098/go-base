package routers

import (
	"HelloBeego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//http://localhost:8080/
	beego.Router("/", &controllers.MainController{})

	//http://localhost:8080/mysql
	beego.Router("/mysql", &controllers.MySqlController{})

	beego.Get("/aa", func(context *context.Context) {
		context.WriteString("我是简单get路由!")
	})

	beego.Post("/bb", func(ctx *context.Context) {
		ctx.WriteString("我是简单post路由!")
	})

	//http://localhost:8080/
	//beego.Router("/api/?:id", &controllers.Demo1Controller{})
	//带？和不带？的区别：带？可以匹配http://localhost:8080/，否则无法匹配
	//beego.Router("/api/:id", &controllers.Demo1Controller{})
	//匹配数字
	//beego.Router("/api/:id([0-9]+)", &controllers.Demo1Controller{})
	//匹配大写字母，小写字母，数字
	//beego.Router("/api/:username([\\w]+)", &controllers.Demo1Controller{})

	//beego.Router("/api/*.*", &controllers.Demo1Controller{})

	//beego.Router(":id:int", &controllers.Demo1Controller{})

	//-----------------------------------自定义方法-----------------------------------------
	//beego.Router("/demo2/aa", &controllers.Demo2Controller{}, "Get:GetFunc;Post:PostFunc")
	//beego.Router("/demo2/aa", &controllers.Demo2Controller{}, "Get,Post:GetAndPostFunc")
	//beego.Router("/demo2/aa", &controllers.Demo2Controller{}, "*:AnyFunc")
	beego.Router("/cookie/aa/xdl", &controllers.Demo2Controller{}, "*:AnyFunc;Post:PostFunc")

	//-----------------------------------自动匹配-------------------------------------------------------------
	//http://localhost:8080/demo3/getfunc/123/456/789
	//http://localhost:8080/demo3/postfunc/123/456/789
	//注意：控制器和方法名必须小写，否则访问不到
	beego.AutoRouter(&controllers.Demo3Controller{})

	//-------------------------------------注解路由----------------------------------------------------------
	//http://localhost:8080/demo4/getfunc/123/456
	beego.Include(&controllers.Demo4Controller{})

	//---------------------------------------cookie-------------------------------------------------
	//   /cookie
	//   /cookie/aa    /cookie/aa/bb   /cookie/dhjsd/dsfld
	beego.Router("/cookie", &controllers.CookieController{})

}
