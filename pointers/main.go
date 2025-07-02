package main

import "fmt"

// modifyValueByReference modifies the value at the given pointer
func modifyValueByReference(num *int) {
	*num = *num * 2 // Doubles the value at the address pointed to by ptr
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num // ptr now holds the address of num

	num := 2
	ptr := &num

	fmt.Println("Value of num:", num)                  // Output: Value of num: 2
	fmt.Println("pointer contains:", ptr)              // Output: pointer ptr: 0x... (address of num)
	fmt.Println("Value of num through pointer:", *ptr) // Output: Value of num through pointer: 2

	var pointer *int //default value of pointer is nil
	if pointer == nil {
		fmt.Println("Pointer is nil")
	}

	value := 5
	modifyValueByReference(&value)
	fmt.Println("Value after modification:", value) // Output: Value after modification: 10
}
