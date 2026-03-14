package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p := person{
		name: "Bob",
		age: 29,
	}
	p.age = 32

	p1 := person{"Robert", 21}
	p2 := person{name: "King"}
	p3 := person{}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}