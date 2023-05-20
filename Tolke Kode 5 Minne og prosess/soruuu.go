package main

import (
	"fmt"
)

func main() {

	var heks [16] byte
	heks = [16] byte {'0', 1,2,3,4,5,6,7,8,9,'A','B','C','D','E','F'}
	fmt.Printf("%v\n", heks)

	var index int = 48
	for j := 0; j < len (heks) -6; j++ {
		heks[j] = byte(index)
		index++
	}

	fmt.Printf ("%s\n", heks)
}
