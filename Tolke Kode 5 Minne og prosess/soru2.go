package main

import "fmt"

func main() {
	var tall uint = 18446744073709551615
	fmt.Println(tall + 1)
	fmt.Println(tall - 1)

	var tall2 uint = 1
	fmt.Println((tall2 << 64) - 1)

}
