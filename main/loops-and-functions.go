package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64  {
	z:=1.0
	j:=1
	for j <100 {
		z-=(z*z-x)/(2*z)
		j+=1
	}
	return z

}
func main() {
	fmt.Print(Sqrt(2),"\n")
	fmt.Print(math.Sqrt(2))
}
