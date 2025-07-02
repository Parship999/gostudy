package main

import "fmt"

func main() {
	// For loop to print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
