package main

import "fmt"

type Set map[interface{}]struct{}

func (s Set) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}
func (s Set) Add(key interface{}) {
	s[key] = struct{}{}
}
func (s Set) Delete(key interface{}) {
	delete(s, key)
}
func main() {
	//map1:=make(map[string]int,10)
	s := make(Set)
	s.Add("Tom")
	s.Add("sam")
	s.Add(1)
	s.Add(3.14)
	//s.Add(map1)
	for k := range s {
		fmt.Printf("%d\n", k)
	}
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Jack"))
}
