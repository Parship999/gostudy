package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Request in Golang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer res.Body.Close() // resource management by closing the connection
	fmt.Printf("Type of res: %T\n", res)

	//read the res body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data: ", string(data)) // convert byte to string
}

/*

In Go, when you read the response body with ioutil.ReadAll(res.Body), you get a value of type []byte (a slice of bytes).

The line:
fmt.Println("Data: ", string(data))

converts the byte slice to a string using string(data). This is necessary because the HTTP response body is raw bytes, but you usually want to print or process it as a human-readable string (for example, JSON text).

*/
