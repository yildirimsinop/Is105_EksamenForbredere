package main

import (
	"fmt"
)

var c, python, java bool
var golang string

func main() {
	var i int
	i = addFive(i)
	fmt.Println(i, c, python, java, golang)
	fmt.Printf("%#b %v %v %v %T\n", i, c, python, java, golang)
}

func addFive(k int) (i int) {
	i = k + 5
	return
}
