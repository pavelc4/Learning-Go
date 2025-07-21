package main

import "fmt"

func main() {

	days := [...]string{"senin", "selasa", "rabu", "jumat", "sabtu", "minggu"}
	daySlice1 := days[4:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur Baru")
	daySlice2[0] = "upps"

	fmt.Println(daySlice2)
	fmt.Println(days)

}
