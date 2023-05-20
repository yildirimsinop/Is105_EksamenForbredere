package main

import (
	"fmt"
)

func main() {

	var heks [16] string
	heks = [16]string {"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}

	var heksTall []string = heks[0:9]
	var heksBokstaver []string = heks[9:]
	fmt.Println(heksTall)
	fmt.Println(heksBokstaver)

	var heksOverlapp []string = heks[8:12]
	fmt.Println(heksOverlapp)

	heksOverlapp[1] = "X"
	fmt.Println(heks)

}
