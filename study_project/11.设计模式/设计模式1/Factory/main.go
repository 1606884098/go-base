package main

import "./Factory"
import "./Abstract_Factory"
import "fmt"

func main1(){
	var fac Factory.OperatorFactory;
	fac=Factory.PlusOperatorFactory{}
	//fac=Factory.SubOperatorFactory{}
	op:=fac.Create()
	op.Setleft(20)
	op.SetRight(10)


	fmt.Println(op.Result())

}
func main(){
	//抽象工厂
	var factory Abstract_Factory.DAOFactory
	//factory=&Abstract_Factory.MySQLFactory{}
	factory=&Abstract_Factory.OracleFactory{}


	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()



}
