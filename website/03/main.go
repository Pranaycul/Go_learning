package main

import "fmt"

const meterToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("enter meters swam :")
	fmt.Scan(&meters)
	yards := meters * meterToYards
	fmt.Println(meters, "meters is equal to  ", yards, " yards.")

}
