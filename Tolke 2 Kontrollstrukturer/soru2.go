package main

import "fmt"

func main() {

	for rad := 1; rad <= 3; rad++ {
		for kol := 1; kol <= 3; kol++ {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}
