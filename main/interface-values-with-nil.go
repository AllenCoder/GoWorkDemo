package main

import "fmt"

type I2 interface {
	M()
}
type T2 struct {
	S string
}

func (t *T) M2()  {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I2
	var t *T2
	i = t
	describe(i)
	i.M()
	i = &T{"hello"}
	describe(i)
	i.M()
}
