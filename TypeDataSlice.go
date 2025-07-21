package main

import "fmt"

func main() {
	Roko := [...]string{"samoerna", "super", "signature", "garpit", "malboro", "magnum"}
	slice1 := Roko[4:6]
	fmt.Println(slice1)

	slice2 := Roko[:2]
	fmt.Println(slice2)

	slice3 := Roko[3:]
	fmt.Println(slice3)

	slice4 := Roko[:]
	fmt.Println(slice4)
}
