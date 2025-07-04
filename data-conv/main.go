package main

import (
	"fmt"
	"strconv" // Importing strconv for string conversion
)

func main() {
	var num int = 42
	fmt.Println("Original number:", num)
	fmt.Printf("type of num is %T\n", num)

	var data float64 = float64(num)
	data += 3.14 // Adding a float to the converted number //data conversion
	// This is a conversion from int to float64
	fmt.Println("Converted number:", data)
	fmt.Printf("type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num) // Convert int to string //itoa stands for "integer to ASCII"
	// This is a conversion from int to string
	fmt.Println("Converted string:", str)
	fmt.Printf("type of str is %T\n", str)

	num_str := "456"
	num_int, _ := strconv.Atoi(num_str) // Convert string to int
	fmt.Println("Converted back to int:", num_int)
	fmt.Printf("type of num_int is %T\n", num_int)

	str_float := "3.14"
	float_num, _ := strconv.ParseFloat(str_float, 64) // Convert string to float
	fmt.Println("Converted string to float:", float_num)
	fmt.Printf("type of float_num is %T\n", float_num)

	//float64 is the default and most commonly used floating-point type in Go because
	//it provides higher precision (about 15 decimal digits) compared to float32 (about 6-7 decimal digits).
}
