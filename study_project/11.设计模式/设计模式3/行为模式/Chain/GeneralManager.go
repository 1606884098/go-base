package Chain

import "fmt"

type GenaralManager struct {
}

func NewGenaralManager() *GenaralManager {
	return &GenaralManager{}
}
func NewGenaralManagerChain() *RequestChain {
	return &RequestChain{&GenaralManager{}, nil}
}

func (pm *GenaralManager) HaveRight(money int) bool {
	return true
}
func (pm *GenaralManager) HandleFeeRequest(name string, money int) bool {
	if name == "weishangyin" {
		fmt.Printf("GenaralManager  授权%s %d\n", name, money)
		return true
	}
	fmt.Printf("GenaralManager not 授权%s %d\n", name, money)
	return false
}
