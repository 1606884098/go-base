package simplemath

import "math"

func Add(a int, b int) int {
	return a + b
}

func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
