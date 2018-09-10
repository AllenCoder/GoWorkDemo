package main

import (
	"fmt"
	"math"
)

type I1 interface {
	M()
}
type T1 struct {
	S string
}

func (t *T1) M1() {
	fmt.Println(t.S)
}

type F1 float64

func (f F1) M() {
	fmt.Println(f)
}
func main() {
	var i I1
	i = &T{"Hello"}
	describe(i)
	i.M()
	i = F1(math.Pi)
	describe(i)
	i.M()
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
