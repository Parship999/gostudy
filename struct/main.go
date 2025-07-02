package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

type contact struct {
	Email string
	Phone string
}

type address struct {
	Street string
	City   string
	State  string
}

type Employee struct {
	Person_Details Person  // embedded struct
	Person_Contact contact // embedded struct
	Person_Address address // embedded struct
}

func main() {
	//1st method // using a variable
	var myself Person
	//fmt.Println("Parship Person:", myself)
	myself.Name = "Parship"
	myself.Surname = "Chowdhury"
	myself.Age = 30
	//fmt.Println("Updated Parship Person:", myself)

	//2nd method // using a struct literal
	person1 := Person{
		Name:    "John",
		Surname: "Doe",
		Age:     25,
	}
	fmt.Println("Person 1:", person1)

	//3rd method //new keyword
	var person2 = new(Person) // its a pointer to Person
	person2.Name = "Jane"
	person2.Surname = "Smith"
	person2.Age = 28
	//fmt.Println("Person 2:", person2.Age)
	// by default the value is 0, // so if i delete the Age field, it will be 0

	var employee1 Employee
	employee1.Person_Details = Person{
		Name:    "Alice",
		Surname: "Johnson",
		Age:     32,
	}

	// employee1.Person_Contact = contact{
	// 	Email: "kdjd.gmail.com",
	// 	Phone: "123-456-7890",
	// }
	employee1.Person_Contact.Email = "dhdh@gmail.com"
	employee1.Person_Contact.Phone = "987-654-3210"

	employee1.Person_Address = address{
		Street: "123 Main St",
		City:   "Springfield",
		State:  "IL",
	}

	fmt.Println("Employee 1 Details:", employee1)
}
