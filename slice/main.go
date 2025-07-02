package main

import (
	"fmt"
)

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 11, 12)

	numbers := make([]int, 4, 10)
	fmt.Println("Original slice:", numbers)
	fmt.Println("Length of slice:", len(numbers))
	fmt.Println("Capacity of slice:", cap(numbers))
}
