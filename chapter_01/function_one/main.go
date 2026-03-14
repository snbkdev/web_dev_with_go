package main

import "fmt"

func pi() float64 {
	return 3.14159
}

func inc(a int) int {
	return a + 1
}

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a := 3
	b := 7

	fmt.Println(pi())
	fmt.Println(inc(a))
	fmt.Println(add(a, b))
}