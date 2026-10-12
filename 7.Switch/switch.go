package main

import "fmt"
import "time"

func main() {
	// Simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// Multiple Conditions switch
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	}

	// Type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I'm an int and my value is %d\n", v)
		case string:
			fmt.Printf("I'm a string and my value is %s\n", v)
		default:
			fmt.Printf("I'm something else and my value is %v\n", v)
		}

	}

	whoAmI("golang")
	whoAmI(3)
	whoAmI(3.14)
}
