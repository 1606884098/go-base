package VD03_dataType01

import (
	cc "container/list"
	"fmt"
	"github.com/1606884098/openSourceTest/mathod"
)

func main() {
	fmt.Printf(mathod.Add())
	heroes := []string{"宋大哥", "吴大哥", "李大哥"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)

	}
	s1 := []int{1, 2, 3, 4}
	s1 = append(s1, 2, 1)
	l := cc.New()
	l.PushBack("1")
	l.PushFront(2)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
