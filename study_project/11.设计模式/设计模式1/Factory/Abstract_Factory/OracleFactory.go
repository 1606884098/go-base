package Abstract_Factory

import "fmt"

//mysql针对两个接口的实现
type OracleMainDAO struct {
}

func (*OracleMainDAO) SaveOrderMain() {
	fmt.Println("Oracle  main save ")
}

type OracleDetailDAO struct {
}

func (*OracleDetailDAO) SaveOrderDetail() {
	fmt.Println("Oracle detail save ")
}
