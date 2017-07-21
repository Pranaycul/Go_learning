package main

import "fmt"

func main() {
	m := make([]string, 1, 23)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Pranay"
	fmt.Println(z)
}
