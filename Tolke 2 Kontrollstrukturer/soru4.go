package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2023, time.January, 20, 23, 59, 59, 0, time.UTC)

	fmt.Println("Når er det Lørdag?")
	today := t.Local().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("I dag.")
	case today + 1:
		fmt.Println("I morgen.")
	case today + 2:
		fmt.Println("Overimorgen")
	default:
		fmt.Println("Lenge til.")
	}

	fmt.Println(t.Local().Day()+1, "Januar", t.Local().Year())
}
