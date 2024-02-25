package main

import "fmt"

func main() {
	a := 10
	b := 2
	addition := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	modulus := a % b

	fmt.Println(a, "+", b, "=", addition)
	fmt.Println(a, "-", b, "=", subtraction)
	fmt.Println(a, "*", b, "=", multiplication)
	fmt.Println(a, "/", b, "=", division)
	fmt.Println(a, "%", b, "=", modulus)

	// Augmented Assignments:
	c := 10
	c += 5 // short way of: c = c + 5
	fmt.Println(c)

	// Unary Operator
	d := 1
	d++ // short way of: d = d + 1
	fmt.Println(d)
	d-- // short way of: d = d - 1
	fmt.Println(d)
}
