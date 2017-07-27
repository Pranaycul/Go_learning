package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type Doublezero struct {
	Person
	First         string
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular Person")
}
func (dz Doublezero) Greeting() {
	fmt.Println("Miss Moneypenny, So good to see you.")
}

func main() {

	p1 := Doublezero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   30,
		},
		First:         "Double zero Seven",
		LicenseToKill: true,
	}
	p2 := Doublezero{
		Person: Person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, " ", p1.Person.First)
	fmt.Println(p2.First, " ", p2.Person.First)

	p1.Greeting()
	p1.Person.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
