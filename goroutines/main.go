package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("World")
}

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	fmt.Println("learning goroutines")
	go sayHello()
	sayHi()

	//wait for goroutine to finish
	time.Sleep(3000 * time.Millisecond)
}
