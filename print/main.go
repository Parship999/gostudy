package main

import "fmt"

func main() {
	age := 25
	name := "alice"
	height := 6.7788888888888888888888

	fmt.Println("age: ", age, "height: ", height)
	fmt.Println("hello")

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %2f\n", name, age, height)
}
