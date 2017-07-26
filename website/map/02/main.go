package main

import "fmt"

func main() {

	var myGreeting = map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}
