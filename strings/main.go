package main

import (
	"fmt"
	"strings"
)

// This program demonstrates the use of the strings package in Go.
func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println("Split parts:", parts)

	str := "Hello, World!"
	count := strings.Count(str, "o")
	fmt.Println("Count of 'o':", count)

	str1 := "        Hello World!  "
	trimmed := strings.TrimSpace(str1)
	fmt.Println("Trimmed string:", trimmed)

	str2 := "Hello"
	str3 := "Universe!"
	result := strings.Join([]string{str2, str3}, ",")
	fmt.Println("Joined string:", result)
}
