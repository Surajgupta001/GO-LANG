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
}