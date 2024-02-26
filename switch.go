package main

import "fmt"

func main() {
	name := "Zainul"

	switch name {
	case "Ajay":
		fmt.Println("Halo, Ajay!")
	case "Zainul":
		fmt.Println("Halo, Zainul!")
	default:
		fmt.Println("May I know u?")
	}
}
