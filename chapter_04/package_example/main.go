package main

import (
	"fmt"
	"web_project/countries"
)

func main() {
	fmt.Println("Let's Go GO!!!")
	fmt.Println("Hello", countries.GetCountry("NW"))
}