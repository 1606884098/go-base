package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["HelloBeego/controllers:Demo4Controller"] = append(beego.GlobalControllerRouter["HelloBeego/controllers:Demo4Controller"],
		beego.ControllerComments{
			Method:           "GetFunc",
			Router:           "/demo4/getfunc/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["HelloBeego/controllers:Demo4Controller"] = append(beego.GlobalControllerRouter["HelloBeego/controllers:Demo4Controller"],
		beego.ControllerComments{
			Method:           "PostFunc",
			Router:           "/demo4/postfunc/:key",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
