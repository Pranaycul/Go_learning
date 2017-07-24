package main

import "fmt"

func main() {
	greatest := max(1000, 20, 30, 44, 33, 277)
	fmt.Println(greatest)
}
func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}

	}
	return largest
}
