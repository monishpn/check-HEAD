package main

import "fmt"

func main() {
	name := "Charlie"
	age := 25
	height := 1.75
	isStudent := true

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Height: %.2f meters\n", height) // .2f formats to 2 decimal places
	fmt.Printf("Is student: %t\n", isStudent)
	fmt.Printf("Name's type: %T\n", name)
	fmt.Printf("Value and Type: %v, %T\n", age, age)
	// this is A
	//this is B
	//this is C
}
