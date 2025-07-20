package main

import "fmt"

func main() {
	var food [3]string

	food[0] = "manggo"
	food[1] = "orange"
	food[2] = "apple"

	fmt.Println(food[0])
	fmt.Println(food[1])
	fmt.Println(food[2])

	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values[0])

}
