package main

import "fmt"

type Vertexxx struct {
	Lat,Long float64
}
var m1 = map[string]Vertexxx{
	"Bell Labs": Vertexxx{
		40.68433, -74.39967,
	},
	"Google": Vertexxx{
		37.42202, -122.08408,
	},
}
func main() {
	fmt.Println(m1)
}
