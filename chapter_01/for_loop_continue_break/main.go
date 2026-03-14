package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < 12; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(a[i])
	}

	for i := 0; i < 12; i++ {
		if i == 3 {
			break
		}
		fmt.Println(a[i])
	}
}