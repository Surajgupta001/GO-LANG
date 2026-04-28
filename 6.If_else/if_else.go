package main

import "fmt"

func main() {

	// if statement => The 'if' statement is used to execute a block of code only if a specified condition is true.
	age := 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("Access granted.")
	}

	// We can declare a variable inside if construct
	if AGE := 15; AGE >= 18 {
		fmt.Println("You are a adult.", AGE)
	} else if AGE >= 13 {
		fmt.Println("You are a teenager.", AGE)
	} else {
		fmt.Println("You are a child.", AGE)
	}

	// NOTE: Go does not have a ternary operator (?:) like some other programming languages.
	// Instead, you can achieve similar functionality using an if-else statement or by using a function.
}