package main

import (
	"fmt"
	"math"
)

var (
	ToBe    bool  = false
	MaxUInt uint8 = 1<<8 - 1
	MaxInt  int8  = 1<<7 - 1
	MinInt  int8  = math.MinInt8
)

func main() {
	ToBe = !ToBe
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxUInt, MaxUInt)
	fmt.Printf("Type: %T Value: %v\n", math.MinInt8, math.MinInt8)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

}
