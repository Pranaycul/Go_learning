package main

import "fmt"

func main() {
	for i := 48; i < 140; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}
