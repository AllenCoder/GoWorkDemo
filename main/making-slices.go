package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlicew("a", a)
	b := make([]int, 0, 5)
	printSlicew("b", b)
	c := b[:2]
	printSlicew("c", c)
	d := c[2:5]
	printSlicew("d", d)

}
func printSlicew(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
