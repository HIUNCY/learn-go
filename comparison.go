package main

import "fmt"

func main() {
	name1 := "Ajay"
	name2 := "Ajay"

	result := name1 == name2 // are name1 equal to name2? returning true
	fmt.Println(result)
	result = name1 != name2 // are name1 not equal to name2? returning false
	fmt.Println(result)

	number1 := 10
	number2 := 7

	result = number1 > number2 // is number1 greater than number2? returning true
	fmt.Println(result)
	result = number1 < number2 // is number1 less than number2? returning false
	fmt.Println(result)
	// note: <= for "less than or equal", and >= for "greater than or equal"
}
