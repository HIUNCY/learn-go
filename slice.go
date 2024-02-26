package main

import "fmt"

func main() {
	days := [7]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	slice1 := days[:] // get all data in array days
	fmt.Println(slice1)
	slice2 := days[4:] // get data start from index 4 to last
	fmt.Println(slice2)
	slice3 := days[:3] // get data start from first to index 2
	fmt.Println(slice3)

	daySlice1 := days[5:] // get data start from index 5 to last
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days) // if u change the value on slice, its also change the original array
}
