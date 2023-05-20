package main

import (
	"fmt"
)

func main() {

	primes := [6] int {2,3,5}
	fmt.Println(primes)

	var sum int
	for _, v := range primes {
		sum += v
	}

	for i := 0 ; i < len(primes); i++ {
		sum += primes [i]
	}

	fmt.Printf("Sum = %d\n", sum)
}
