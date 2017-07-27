package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
