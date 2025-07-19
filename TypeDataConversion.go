package main

import "fmt"

func main() {

	var Niali32 int32 = 32768
	var Nilai64 int64 = int64(Niali32)
	var Nilai16 int16 = int16(Niali32)

	fmt.Println(Niali32)
	fmt.Println(Nilai64)
	fmt.Println(Nilai16)

	var animal = "Cat"
	var a = animal[0]
	var aString = string(a)

	fmt.Println(animal)
	fmt.Println(aString)

}
