package main

import "fmt"

func main () {
	// var name string = "Go Lang"

	// Infer => Go can infer the type of the variable based on the value assigned to it.
	var name = "Go Lang"

	var isAdult = true // boolean variable to indicate if someone is an adult or not.

	var age int = 25 // integer variable to store the age of a person.

	// SHORT HAND SYNTAX => Go provides a shorthand syntax for declaring and initializing variables.
	// This syntax is only available inside functions and cannot be used at the package level.
	// The syntax is: variableName := value
	// The type of the variable will be inferred from the value assigned to it.
	// name := "Go Lang"
	// isAdult := true
	// age := 25

	var height string
	height = "6 feet"

	// var price float32 = 50.5
	// price = 50.5 // This will cause an error because the variable 'price' is already declared in the same scope.
	price := 50.5 // This is the correct way to declare and initialize the variable 'price' using the shorthand syntax.

	fmt.Println(name)
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(price)
}