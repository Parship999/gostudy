package main

import (
	"encoding/json"
	"fmt"
)

// This is a Go struct definition for a type named person.
type person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"` //agar is_adult ye wala data nahi hai toh by default IsAdult use hoga
}

func main() {
	fmt.Println("Hello World")
	p := person{Name: "John", Age: 20, IsAdult: true}
	fmt.Println("person data : ", p)

	//convert person struct to json encoding (marshalling)
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("json data : ", string(jsonData)) //convert byte to string because json is byte array

	//convert json to person struct decoding (unmarshalling)
	var p2 person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("person data : ", p2)
}
