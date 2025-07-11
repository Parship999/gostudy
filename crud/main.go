package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"` //Field name in Go can be anything (Compl, Completed, etc.).JSON tag (json:"completed") controls mapping between JSON and struct field.
}

func main() {
	fmt.Println("Hello World")
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
