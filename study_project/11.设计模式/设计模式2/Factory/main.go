package main

import "./Factory"
import "./Abstract_Factory"
import "./Builder"
import "./Composite"
import "./Bridge"
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
func main2(){
	//抽象工厂
	var factory Abstract_Factory.DAOFactory
	//factory=&Abstract_Factory.MySQLFactory{}
	factory=&Abstract_Factory.OracleFactory{}


	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()



}
func main3(){
	//builder:=&Builder.StringBuilder{}
	builder:=&Builder.IntBuilder{}
	dict:=Builder.NewDirector(builder)
	dict.Makedata()
	fmt.Println(builder.GetResult())
}
func main4(){
	//root;
	root:=Composite.NewComonent(Composite.CompositeNode,"root")
	r1:=Composite.NewComonent(Composite.CompositeNode,"r1")
	r2:=Composite.NewComonent(Composite.CompositeNode,"r2")
	r3:=Composite.NewComonent(Composite.CompositeNode,"r3")

	l1:=Composite.NewComonent(Composite.LeafNode,"l1")
	l2:=Composite.NewComonent(Composite.LeafNode,"l2")
	l3:=Composite.NewComonent(Composite.LeafNode,"l3")

	root.AddChild(r1)
	root.AddChild(r2)
	r1.AddChild(r3)

	r1.AddChild(l1)

	r2.AddChild(l2)
	r2.AddChild(l3)
	root.Print("")


}
func main(){
	//m:=Bridge.NewComonMessage(Bridge.ViaSMS())
	//m:=Bridge.NewComonMessage(Bridge.ViaEmail())
	//m.SendMessage("baBy 你好","nimei")
	m:=Bridge.NewUrencyMessage(Bridge.ViaEmail())
	m.SendMessage("baBy 你好","nimei")
}