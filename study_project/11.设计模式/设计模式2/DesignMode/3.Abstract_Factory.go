package main

//sqlsever
//mysql      sql
//oracle

//订单
//订单报表

type OrderMainDAO interface { //订单记录
	SaveOrderMain() //保存，
	//DeleleOrderMain()
	//SearchOrderMain()
}
type OrderDetailDAO interface { //订单详情
	SaveOrderDetail() //保存
}
type DAOFactory interface { //抽象工厂接口
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//详情请见多文件抽象工厂
