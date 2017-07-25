package main

import "fmt"

func main() {

	recored := make([][]string, 0)

	var student1 = make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100%"
	student1[3] = "74.45%"

	recored = append(recored, student1)

	var student2 = make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "94%"
	student2[3] = "96%"

	recored = append(recored, student2)

	fmt.Println(recored)

}
