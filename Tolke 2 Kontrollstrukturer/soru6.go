package main

import "fmt"

func printNumber(i int) {
	defer fmt.Println(i)
}

func main() {
	fmt.Println("teller")

	defer fmt.Println("ferdig")

	for i := 0; i < 3; i++ {
		defer printNumber(i)
		defer printNumber(i * i)
	}
}
