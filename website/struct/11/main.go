package main

import (
	"encoding/json"
	"fmt"

	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	var p1 Person
	rds := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rds).Decode(&p1)
	fmt.Println("---------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}
