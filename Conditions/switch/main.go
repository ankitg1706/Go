package main

import "fmt"

func main() {

	input := 6

	switch input {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Holiday")
	}
}
