package main

import "fmt"

func main() {
	name1 := "Ajay"
	name2 := "Ajay"

	result1 := name1 == name2 // are name1 equal to name2? returning true
	fmt.Println(result1)
	result2 := name1 != name2 // are name1 not equal to name2? returning false
	fmt.Println(result2)

	number1 := 10
	number2 := 7

	result3 := number1 > number2 // is number1 greater than number2? returning true
	fmt.Println(result3)
	result4 := number1 < number2 // is number1 less than number2? returning false
	fmt.Println(result4)
	// note: <= for "less than or equal", and >= for "greater than or equal"
}
