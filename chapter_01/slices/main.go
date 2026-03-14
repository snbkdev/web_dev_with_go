package main

import "fmt"

func main() {
	a := [10]int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	
	// s := a[3:8]
	// s = make([]int, 10)
	// var s []int = a[1:6]
	
	var s1 []int = a[0:5]
	var s2 []int = a[:5]
	var s3 []int = a[0:]
	var s4 []int = a[:]

	s := []int{1, 4, 9, 16, 25}

	s = append(s, 18, 27)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println(s)
}