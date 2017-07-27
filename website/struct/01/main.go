package main

import "fmt"

func main() {

	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{"Pranay", "Wankhede", 29}
	p2 := person{first: "Phoenix", last: "Ash-maker", age: 0}

	fmt.Println(p1.first, " ", p1.last, " ", p1.age)
	fmt.Println(p2.first, " ", p2.last, " ", p2.age)

}
