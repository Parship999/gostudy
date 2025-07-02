package main

import "fmt"

func main() {
	// Create a map to store the names and ages of people
	peopleAge := make(map[string]int)

	// Add some entries to the map
	peopleAge["Alice"] = 30
	peopleAge["Bob"] = 25
	peopleAge["Charlie"] = 35

	fmt.Println("bob :", peopleAge["Bob"]) // Accessing a value by key
	peopleAge["Bob"] = 26                  // Update the value for Bob
	fmt.Println("Updated Bob's age:", peopleAge["Bob"])

	// Remove Charlie from the map
	delete(peopleAge, "Charlie")
	fmt.Println("After deleting Charlie:", peopleAge)

	// Check if Alice exists in the map
	// map returns two values: the value and a boolean indicating if the key exists
	People, Exists := peopleAge["Parship"]
	fmt.Println("Parship's age:", People) // This will print 0 if Parship does not exist
	fmt.Println("Does Parship exist?:", Exists)

	// Iterate over the map and print each name and age
	for name, age := range peopleAge {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
