package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ {
		div := func(i int) (int, bool) {

			return i / 2, i%2 == 0
		}

		h, even := div(x)

		fmt.Println(h, even)
	}
}
