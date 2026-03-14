package main

import "fmt"

func double(a int) {
	a = a * 2
	fmt.Println(a)
}

func doublePointer(a *int) {
	*a = *a * 2
	fmt.Println(*a)
}

func modify(s []int) {
	s[0] = 4
	s = append(s, 5)
	fmt.Println(s)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	modify(s)
	fmt.Println(s)
}

func calc() func() int {
	a := 0
	return func() int {
		a += 1
		return a
	}
}