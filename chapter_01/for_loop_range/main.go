package main

import "fmt"

func main() {
	var arr [20]int
	for i := 0; i < 20; i++ {
		arr[i] = i * 2
	}

	sum := 0
	for i, v := range arr {
		if i%2 == 0 {
			sum += v
		}
	}
	fmt.Println(sum)
}