package main

import "fmt"

func zero(x *int) {
	fmt.Println("value of x  is ", x)
	fmt.Println("address of x is ", &x)
	*x = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println("Value of x is ", x)
	fmt.Println("address of x is ", &x)
}
