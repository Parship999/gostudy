package main

import "fmt"

func main() {
	// Example 1: Simple if-else statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// Example 2: If-else with multiple conditions
	y := 20
	if y < 10 {
		fmt.Println("y is less than 10")
	} else if y < 30 {
		fmt.Println("y is between 10 and 30")
	} else {
		fmt.Println("y is 30 or more")
	}
}
