package main

import "fmt"

func main() {
	// Defer is used to ensure that a function call is performed later in a program's execution, usually for cleanup purposes.
	// It is often used to close files, release resources, or unlock mutexes.

	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")

	// Defer statements are executed in LIFO (Last In, First Out) order.
	defer fmt.Println("This will be printed second")
	defer fmt.Println("This will be printed third")
}
