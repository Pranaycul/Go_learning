package main

import "fmt"

func main() {
	records := make([][]int, 0)

	for i := 0; i < 3; i++ {
		record := make([]int, 0)

		for j := 0; j < 4; j++ {

			record = append(record, j)
		}
		records = append(records, record)
	}

	fmt.Println(records)

}
