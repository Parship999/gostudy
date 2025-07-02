package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey!!!!")

	// var name string;
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name)
}

//scan can't read after 1st space
//so we have to use the buffered reader