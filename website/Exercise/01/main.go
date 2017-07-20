package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBizz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Bizz")
			continue
		} else {
			fmt.Println(i)
		}

	}

}
