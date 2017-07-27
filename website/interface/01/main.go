package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}
func main() {
	s := Square{10}
	info(s)
}
