package main

import (
	"flag"
	"fmt"
)

var name string

var sex int

func init() {
	flag.IntVar(	&sex ,"sex", 1, "The sex")
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}
func main() {
	flag.Parse()
	fmt.Printf("Hello, name=%s sex=%s!\n",name, sex)
}
