package main

import "fmt"

func main(){
	var name[5]string
	name[0] = "prince"

	fmt.Println("Name: ", name)

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))


	var name1[5] string
	name1[2] = "Parship"
	name1[0] = "Akash"

	fmt.Println(name1)
	fmt.Printf("%q\n",name1)
}