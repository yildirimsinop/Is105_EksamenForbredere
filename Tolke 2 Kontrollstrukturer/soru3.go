package main

import "fmt"

func openDoor (s bool) bool {
	if s {
		fmt.Println("Door already open.")
	} else {
		s = !s
		fmt.Println("Door is now open.")
	}

	return s
}

func closeDoor (s bool) bool {
	if !s {
		fmt.Println("Door already closed.")
	} else {
		s = !s
		fmt.Println("Door is now closed.")
	}

	return s
}

func main () {
	var stateOfTheDoor bool = false
	stateOfTheDoor = openDoor(stateOfTheDoor)
	stateOfTheDoor = closeDoor(stateOfTheDoor)
	closeDoor(stateOfTheDoor)
}
