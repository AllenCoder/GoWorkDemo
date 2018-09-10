package main

import (
	"math"
	"fmt"
)

type Vertext7 struct {
	X, Y float64
}

func (v *Vertext7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f

}
func (v Vertext7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}
func main() {
	v := &Vertext7{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

}
