package main

import (
	"fmt"
	"math"
)

func main() {
	var minFloat64 float64 = math.SmallestNonzeroFloat64
	var maxFloat64 float64 = math.MaxFloat64
	var minFloat32 float32 = math.SmallestNonzeroFloat32
	var maxFloat32 float32 = math.MaxFloat32

	fmt.Println(maxFloat64 / -minFloat64)
	fmt.Println(minFloat64 / -maxFloat64)
	fmt.Println(maxFloat32 / minFloat32)
	fmt.Println(minFloat32 / maxFloat32)
}
