package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type MySqlController struct {
	beego.Controller
}

/*
sqluser=root
sqlpassword=111111
sqlport=3306
sqlurls=127.0.0.1
*/
func (this *MySqlController) Get() {
	sqluser := beego.AppConfig.String("sqluser")
	sqlpassword := beego.AppConfig.String("sqlpassword")
	/*	sqlport, err := beego.AppConfig.Int("sqlport")
		if err != nil {
			sqlport = 3306
		}*/
	sqlport := beego.AppConfig.DefaultInt("sqlport", 3306)

	sqlurls := beego.AppConfig.String("sqlurls")
	this.Ctx.WriteString("sqluser = " + sqluser + "\n" + "sqlpassword = " + sqlpassword +
		"\nsqlport = " + strconv.Itoa(sqlport) + "\nsqlurls = " + sqlurls)
}
