package main

import "fmt"

func main() {

	type PhoneNumber string
	var customerPhone PhoneNumber = "08923623232"

	var rawString string = "086273681236"
	var convertedPhone PhoneNumber = PhoneNumber(rawString)

	fmt.Println("Customer Phone:", customerPhone)
	fmt.Println("Converted Phone:", convertedPhone)
}
