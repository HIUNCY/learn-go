package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80 // returning true
	lulusAbsensi := absensi > 80       // returning false

	andOperator := lulusNilaiAkhir && lulusAbsensi // Returns true if all statements are true, returning false
	fmt.Println(andOperator)

	orOperator := lulusNilaiAkhir || lulusAbsensi // Returns true if one of the statements are true, returning true
	fmt.Println(orOperator)

	notOperator := !andOperator // Reverse the result, returning true cause the result is false
	fmt.Println(notOperator)
	notOperator = !orOperator // Reverse the result, returning false cause the result is true
	fmt.Println(notOperator)

}
