package main

import (
	"fmt"
)

func main() {

	var a [2]string = [2]string{"You", "too"}
	a[1] = "Hello"
	a[0] = "World"
	fmt.Println(a)
}
