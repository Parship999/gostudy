package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// Format the current time
	formattedTime := currentTime.Format("02-01-2006, Monday, 3:04 PM") //golang does not support dd-mm-yyyy format directly //it uses 02-01-2006 for date formatting, Monday for day of the week, and 3:04 PM for time
	fmt.Println("Formatted Time:", formattedTime)

	// Parse a date string into a time.Time object
	layout_string := "2006-01-02" // This layout is used to parse dates in the format YYYY-MM-DD
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_string, dateStr) // _ is used to ignore the error for simplicity
	fmt.Println("Formatted Time:", formatted_time)

	//add 1 day to the current time
	oneDayLater := currentTime.Add(24 * time.Hour)
	fmt.Println("One Day Later:", oneDayLater)
	formatted_new_data := oneDayLater.Format("02-01-2006, Monday, 3:04 PM")
	fmt.Println("Formatted Time After Adding One Day:", formatted_new_data)
}

//The function time.Parse returns two values: the parsed time and an error.
//The function oneDayLater.Format formats the time in a specific layout.
