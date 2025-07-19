package main

import "fmt"

func main() {
	var (
		a = 10
		b = 5

		sum        = a + b
		difference = a - b
		product    = a * b
		quotient   = a / b
	)

	fmt.Println("Penjumlahan:", sum)
	fmt.Println("Pengurangan:", difference)
	fmt.Println("Perkalian:", product)
	fmt.Println("Pembagian:", quotient)

	//Augmented Assignments
	var i = 10
	i += 30
	fmt.Println(i)

	i -= 5
	fmt.Println(i)

	i *= 5
	fmt.Println(i)

	i /= 5
	fmt.Println(i)

	i %= 5
	fmt.Println(i)
}
