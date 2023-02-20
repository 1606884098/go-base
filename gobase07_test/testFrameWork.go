package demo

import "net/http"

func GetArea(weight int, height int) int {
	return weight * height
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
