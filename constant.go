package main

import "fmt"

func main() {
	const name = "Ajay"
	fmt.Println(name) // success

	name = "Zainul"
	fmt.Println(name) // error: cannot reassign const variable

	const (
		firstName = "Muhamad"
		midName   = "Zainul"
		lastName  = "kamal"
		nickName  = "Ajay"
	) // multiple const variable
	fmt.Println(firstName, midName, lastName)
	fmt.Println(nickName)
}
