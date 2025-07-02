package main

import "fmt"

func main() {
	// Switch statement to determine the day of the week
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch statement to determine the quarter of the year
	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("It's the first quarter of the year.")
	case "April", "May", "June":
		fmt.Println("It's the second quarter of the year.")
	case "July", "August", "September":
		fmt.Println("It's the third quarter of the year.")
	case "October", "November", "December":
		fmt.Println("It's the fourth quarter of the year.")
	default:
		fmt.Println("Invalid month")
	}

	// Switch statement to determine the temperature range
	temperature := 30

	switch {
	case temperature < 0:
		fmt.Println("It's freezing!")
	case temperature >= 0 && temperature < 20:
		fmt.Println("It's cold.")
	case temperature >= 20 && temperature < 30:
		fmt.Println("It's warm.")
	case temperature >= 30:
		fmt.Println("It's hot.")
	default:
		fmt.Println("Temperature is not defined.")
	}
}
