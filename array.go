package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "yayat"
	name[1] = "yaya"
	name[2] = "kons"

	fmt.Println(name[0])
	fmt.Println(name)

	var values = [...]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(cap(values))
	fmt.Println(values)

}
