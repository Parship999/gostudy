package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("we are learning function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of 2 number is:", ans)
}
