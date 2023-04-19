package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("%T %T %T %T %T\n", i, j, c, python, java)
	fmt.Println(board())
	// fmt.Printf ("%v/n", board ())
}

func board() (c00, c01, c02, c10, c11, c12, c20, c21 string, c22 int) {
	c00, c01, c02 = "x", "o", "x"
	c10, c11, c12 = "o", "0", "x"
	c20 = "o"
	c21 = "x"
	c22 = 0
	return
}
