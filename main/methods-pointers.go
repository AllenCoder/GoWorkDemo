package main

import (
	"math"
	"fmt"
)

type Vertex4 struct {
	X, Y float64
}

func Abs1(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}
func Scale(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex4{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs1(v))

}
