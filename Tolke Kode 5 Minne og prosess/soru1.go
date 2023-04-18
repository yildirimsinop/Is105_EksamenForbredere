package main

import "fmt"

func main() {
	var tall int16 = -32768

	fmt.Println(tall + 1) // 32767
	fmt.Println(tall - 1) // -32769
}
