package main

import "fmt"

func main() {
	a := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < 20; i++ {
		a[i] = i * 2
	}

	fmt.Println(a)

	result := 10
	sum := 5
	for result < 500 {
		result *= sum*2
	}

	fmt.Println(result)
}