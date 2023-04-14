package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("God morgen!")
	case t.Hour() < 17:
		fmt.Println("God ettermiddag.")
	default:
		fmt.Println("God kveld.")
	}
}
