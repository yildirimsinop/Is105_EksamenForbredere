package main

import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
	name string
}

func main() {

	fmt.Println(Vertex {1, 2, "Punkt A"})
	v := Vertex {1, 5, "Punkt A"}
	fmt.Printf("%v %v %q\n", v.X, v.Y, v.name)

	var v1 * Vertex = new (Vertex)
	v1.X = 4
	v1.Y = 8
	v1.name = "Punkt B"
	fmt.Printf("%T\n", *v1)
	fmt.Println(v1)
	
	v2 := Vertex {} 
	v2.X = 23
	v2.Y = 24
	v2.name = "Punkt C"
	fmt.Printf("%T\n", v2)
	fmt.Println(v2)
	}
