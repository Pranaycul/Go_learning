package main

import "fmt"

func main() {
	var x [258]byte

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i <= 257; i++ {
		x[i] = byte(i)
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
