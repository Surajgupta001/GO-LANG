package main

import "fmt"

// for -> The 'for' loop is the only loop construct in Go.
// It can be used to iterate over a range of values, execute a block of code a specific number of times, or loop indefinitely until a certain condition is met.

func main() {
	// while loop => Go does not have a 'while' loop, but you can achieve the same functionality using a 'for' loop with a condition.
	// Example of a while loop using a for loop:
	count := 0
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	// Infinite loop => You can create an infinite loop using a 'for' loop without any conditions.
	// for {
	// 	fmt.Println("This is an infinite loop. Press Ctrl+C to stop it.")
	// }

	// Classic for loop => The classic for loop has three components: initialization, condition, and post statement.
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)

		if i == 2 {
			continue // Skip the rest of the loop body and move to the next iteration when i is 2.
		}

		fmt.Println("This will not be printed when i is 2.")
	}

	// Range-based for loop => The 'range' keyword is used to iterate over elements in a collection such as arrays, slices, maps, or strings.
	for i := range 6 { // Here we are using 'range' to iterate over a range of values from 0 to 5.
		fmt.Println("Index:", i)
	}
}