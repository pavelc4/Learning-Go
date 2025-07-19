package main

import "fmt"

func main() {
	var NinailAhir = 90
	var Absensi = 90

	var LulusNilaiAhir bool = NinailAhir > 80
	var LulusAbsensi bool = Absensi > 80

	var Lulus bool = LulusNilaiAhir && LulusAbsensi

	fmt.Println(Lulus)
}
