package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	//"io/ioutil"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"` //Field name in Go can be anything (Compl, Completed, etc.).JSON tag (json:"completed") controls mapping between JSON and struct field.
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/4")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status code error: ", res.StatusCode)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Todo: ", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "New Todo",
		Completed: true,
	}

	//convert todo struct to json //web request mein jab bhi data transfer karte hai to json mein hota hai
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert json to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	//perform post request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Updated Todo parship",
		Completed: false,
	}

	//convert todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//convert json to string
	jsonString := string(jsonData)
	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	//perform put request
	req, err := http.NewRequest("PUT", myURL, jsonReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")

	//perform request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

	//check status code
	fmt.Println("Status code error: ", res.Status)
}

func performDeleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//perform delete request
	req, err := http.NewRequest("DELETE", myURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//perform request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	//check status code
	fmt.Println("Status code error: ", res.Status)
}

func main() {
	fmt.Println("Hello World")
	//performGetRequest()
	//performPostRequest()
	//performUpdateRequest()
	performDeleteRequest()
}

/*

If you change both the field name and the JSON tag, for example:

then Go will expect the JSON data to have a key named "compl" instead of "completed".
If the JSON from the API still uses "completed", the field Compl will not be set (it will be false by default for bool).

Summary:

If the JSON tag does not match the key in the JSON data, the field will not be populated.
The field name in Go can be anything, but the JSON tag must match the JSON key for correct mapping.

*/

/*

If you change both the field name and the JSON tag for Title, for example:

Then Go will look for a key named "heading" in the JSON data.
If the JSON from the API still uses "title", the Heading field will not be set (it will be an empty string by default).

Summary:

The struct field name (Heading) is for your Go code.
The JSON tag (json:"heading") must match the key in the JSON data for the field to be populated.
If the JSON key does not match, the field will have its zero value (empty string for string).

*/

/*

The line req.Header.Set("Content-Type", "application/json") is important because:

It tells the server what type of data we're sending in the request body
For REST APIs, application/json is the standard content type for JSON data
Without this header, the server might not know how to parse the request body correctly
Here's how it works in your code:

// First you create JSON data
jsonData, err := json.Marshal(todo)

// Then when sending the request, you tell the server it's JSON
req.Header.Set("Content-Type", "application/json")

Think of it like putting a label on a package - the Content-Type header tells the server "this package contains JSON data, please handle it accordingly".
If you didn't set this header:

The server might reject the request
The server might not parse the JSON data correctly
The server might treat the data as plain text

*/
