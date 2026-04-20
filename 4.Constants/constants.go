package main

import "fmt"

const age = 30 // Outside of the main function, constants are declared at the package level and can be accessed throughout the package.

func main() {
	// Constants are immutable values that cannot be changed after they are declared.
	// They are declared using the 'const' keyword followed by the name of the constant and its value.
	const Pi = 3.14
	const Greeting = "Hello, World!"
	const IsGoFun = true

	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Greeting:", Greeting)
	fmt.Println("Is Go fun?", IsGoFun)

	// Constants can also be declared using a block of constants.
	const (
		SpeedOfLight = 299792458 // in meters per second
		GravitationalConstant = 6.67430e-11 // in m^3 kg^-1 s^-2
		PlanckConstant = 6.62607015e-34 // in J s
	)

	fmt.Println("Speed of Light:", SpeedOfLight, "m/s")
	fmt.Println("Gravitational Constant:", GravitationalConstant, "m^3 kg^-1 s^-2")
	fmt.Println("Planck Constant:", PlanckConstant, "J s")

	// Constants can also be used in expressions and calculations.
	const Radius = 5.0
	const Area = Pi * Radius * Radius
	fmt.Println("Area of the circle with radius", Radius, "is", Area)
}