package main

import "fmt"

func main() {
	const Food = "meatball"
	const Car = "Hyundai"

	// This will print the constant values
	fmt.Println("Favorite food:", Food)
	fmt.Println("Car brand:", Car)

	// The lines below are invalid in Go and will cause a compile error
	// Food = "bakso" // ❌ cannot assign to Food (it's a constant)
	// Car = "audi"   // ❌ cannot assign to Car (it's a constant)

	const (
		animal = "Camel"
		name   = "chiken"
	)

	fmt.Println(animal)
	fmt.Println(name)
}
