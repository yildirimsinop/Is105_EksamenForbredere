package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 32
	Small = Big >> 31
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Printf("%.f\n", needFloat(Big))
	fmt.Printf("%.f\n", math.Pow(2, 22))
	fmt.Printf("%#08b , %d\n", 1, 1)
	fmt.Printf("%#08b, %d\n", 1<<1, 1<<1)
	fmt.Printf("%#08b, %d\n", 1<<2, 1<<2)
	fmt.Printf("%#08b, %d\n", 2<<3, 2<<3)
	fmt.Printf("%#08b, %d\n", 2>>100, 2>>100)
}
