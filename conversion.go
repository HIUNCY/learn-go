package main

import "fmt"

func main() {
	var number32 int32 = 32767
	var number64 int64 = int64(number32)
	var number16 int16 = int16(number32)

	fmt.Println(number32)
	fmt.Println(number64)
	fmt.Println(number16)

	name := "Ajay"
	var j uint8 = name[1] // still byte
	jString := string(j)  // conversing the byte of "j" to string

	fmt.Println(name)
	fmt.Println(j)
	fmt.Println(jString)
}
