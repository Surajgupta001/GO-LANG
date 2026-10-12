package main

import "fmt"

// Numbered squence of specific length. Arrays are used to store multiple values of the same type in a single variable.
func main() {
	var nums [10]int // Declare an array of integers with a length of 10

	nums[0] = 1 // Assign values to the array
	fmt.Println(nums[0])

	fmt.Println(nums) // Print the entire array

	// Array length
	fmt.Println(len(nums))

	// Initialize an array of booleans
	var vals[4]bool
	vals[0] = true
	fmt.Println(vals)

	var names[3]string
	names[0] = "Alice"
	fmt.Println(names)

	// IMP NOTE : int -> 0, string -> "", bool -> false

	// Declaration and initialization of an array
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 2d array
	var arr2d = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr2d)

	// IMPORTANT POINT:
	// -> fixed size, that is predictable and known at compile time.
	// -> memory is allocated on the stack, which can lead to better performance for small arrays.
	// -> cannot time access elements, as the size is fixed and known at compile time.
}