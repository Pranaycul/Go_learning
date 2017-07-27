package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {

	p1 := person{"Pranay", "Wankhede", 29}
	p2 := person{first: "Phoenix", last: "Ash-maker", age: 0}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}
