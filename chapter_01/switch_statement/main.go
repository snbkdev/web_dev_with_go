package main

import "fmt"

func main() {
	code := "fr"
	var country string
	switch code {
	case "fr":
		country = "France"
	case "uk":
		country = "United Kingdom"
	default:
		country = "China"
	}

	fmt.Println(country)

	number := 7
	switch {
	case number % 2 == 0:
		fmt.Println("Even number")
	case number % 2 == 1:
		fmt.Println("Odd Number")
	default:
		fmt.Println("Invalid number")
	}
}