package main

import "fmt"

func main() {
	sum := 0
	for i := 5; i > 0; i-- {
		sum += i
	}

	fmt.Println(sum)
}
