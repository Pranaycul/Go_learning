package main

import "fmt"

func filler(number []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range number {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filler([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)

}
