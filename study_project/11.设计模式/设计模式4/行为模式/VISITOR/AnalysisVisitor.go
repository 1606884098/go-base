package VISITOR

import "fmt"

type AnalysisVisitor struct {
}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("Analysis  Enterprise  customer", c.name)
	}
}
