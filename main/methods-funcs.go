package main

import (
	"math"
	"fmt"
)

type Vettex1 struct {
	X, Y float64
}

func Abs(v Vettex1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vettex1{3, 4}
	fmt.Println(Abs(v))
}
