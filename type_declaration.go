package main

import "fmt"

func main() {
	type noKTP string // noKTP is alias for string

	var ktpAjay noKTP = "11111111111"

	contoh := "22222222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpAjay)
	fmt.Println(contohKTP)
}
